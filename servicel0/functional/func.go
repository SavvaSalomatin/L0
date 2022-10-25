package functional

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"os"
	"os/signal"
	"servicel0/config"
	http_server "servicel0/http-server"
	"time"

	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
)

type Service struct{}

func Docache(db *sql.DB) error {

	rows, err := db.Query("SELECT * FROM schema.l0")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	for rows.Next() {
		var order_uid string
		var data []byte
		rows.Scan(&order_uid, &data)
		order := config.Order{order_uid, data}
		config.C.Set(order_uid, order, cache.DefaultExpiration)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
	} else {
		log.Println("Кеш успешно восстановлен из БД")
	}
	return err
}

func Validcheck(f []byte) bool {
	var validating *validator.Validate
	validating = validator.New()
	var checktype config.Validdata
	err := json.Unmarshal(f, &checktype)
	if err != nil {
		log.Println("Ошибка типов данных валидируемых значений")
		log.Println(err)
		return false
	}
	err = validating.Struct(checktype)
	if err != nil {
		log.Println(err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return false
	}
	return true
}

func StanConnect(cluster_id string, client_id string) (stan.Conn, time.Duration) {
	sc, _ := stan.Connect(cluster_id, client_id, stan.Pings(5, 2), stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
		log.Fatalf("Связь с nats-сервером потеряна, причина: %v", reason)
	}))
	aw, _ := time.ParseDuration("10s")
	log.Println("Успешное подключение к nats-streaming-серверу")
	return sc, aw
}

func Subscribee(db *sql.DB, theme *string, cluster_id *string, client_id *string) {
	sc, aw := StanConnect(*cluster_id, *client_id)
	sub, err := (sc).Subscribe(*theme, func(m *stan.Msg) {
		m.Ack()
		datawork(m, db)
	}, stan.SetManualAckMode(), stan.AckWait(aw), stan.DurableName(*client_id))
	if err != nil {
		log.Println(err)
	}
	listen(&sub, &sc)
}

func datawork(m *stan.Msg, db *sql.DB) {
	var data map[string]interface{}
	log.Printf("Получено из канала: %s\n", string(m.Data))
	json.Unmarshal(m.Data, &data)
	if !Validcheck(m.Data) {
		log.Printf("Получены некорректные данные")
	} else {
		foo, order_uid := dataCacheSet(&data)
		addtodb(db, foo, order_uid)
	}
}

func dataCacheSet(data *map[string]interface{}) (*[]byte, *string) {
	uid := (*data)["order_uid"].(string)
	_, found := config.C.Get(uid)
	delete(*data, "order_uid")
	foo, _ := json.Marshal(data)
	order := config.Order{uid, foo}
	if found {
		log.Println("Данные с указанным order_uid уже находятся в системе")
		return &foo, &uid
	} else {
		config.C.Set(uid, order, cache.DefaultExpiration)
		log.Println("Данные добавлены в кеш")
		return &foo, &uid
	}
}

func addtodb(db *sql.DB, foo *[]byte, order_uidbuf *string) {
	var ErrUnique = errors.New("pq: повторяющееся значение ключа нарушает ограничение уникальности \"unique_uid\"")
	result, err := db.Exec("insert into schema.l0 (order_uid, jsondata) values ($1, $2)",
		*order_uidbuf, *foo)
	if err != nil {
		if err.Error() == ErrUnique.Error() {
			log.Println("Повторение уникального значения" + *order_uidbuf + ", данные не были добавлены")
		} else {
			log.Println(err)
			log.Println(result.RowsAffected())
		}
	} else {
		log.Println("Полученное сообщение успешно добавлено в БД")
	}

}

func listen(sub *stan.Subscription, sc *stan.Conn) {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			log.Printf("\nСработало прерывание. Закрытие соединения...\n\n")
			(*sub).Unsubscribe()
			(*sc).Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}

func Connectdb(user *string, password *string, dBName *string) *sql.DB {
	var confstrpostgres = "user=" + *user + " password=" + *password + " dbname=" + *dBName + " sslmode=disable"
	db, err := sql.Open("postgres", confstrpostgres)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Успешное подключение к БД")
	}
	return db
}

func (service Service) Start(clusterID *string, clientID *string, theme *string, user *string, password *string, dbName *string, port *string) {
	db := Connectdb(user, password, dbName)
	Docache(db)
	server := http_server.Httpsrv{}
	go Subscribee(db, theme, clusterID, clientID)
	go server.Httpstart(port)
	defer db.Close()
	fmt.Scanln()
}
