package main

import (
	"flag"
	"fmt"
	"log"
	"servicel0/functional"
)

var usageStr = `
Использование: main.go [обязательные опции]
_________________________________________________________________________________________________
Пример использования:
go run main.go -cl_id test-cluster -cli_id clientid -t fee -u postgres -p 123 -db TestL -prt 8080
_________________________________________________________________________________________________
Опции:
	-cl_id,  		--cluster  <cluster name>           ID кластера
	-cli_id,  		--client  <client ID>  				ID клиента(сервиса)
	-t, 			--theme  <theme>           			Тема подписки
	-u,  			--user  <user>             			Имя пользователя СУБД
	-p, 			--password    <password>    		Пароль пользователя СУБД
	-db,  			--database_name  <database name>   	Имя подключаемой БД
	-prt, 			--http_port <http port>             Порт работы http-сервера
_________________________________________________________________________________________________
Невведенные опции во время работы программы заменяются на стандартные 
		ClusterID test-cluster
		ClientID  clientidid
		Theme     foo
		User      postgres
		Password  123
		DbName    TestL
		HTTPPort  8080
__________________________________________________________________________________________________
`

func main() {
	var (
		ClusterID string
		ClientID  string
		Theme     string
		User      string
		Password  string
		DbName    string
		HTTPPort  string
	)
	flag.StringVar(&ClusterID, "cl_id", "test-cluster", "ID кластера")
	flag.StringVar(&ClientID, "cli_id", "clientidid", "ID клиента(сервиса)")
	flag.StringVar(&Theme, "t", "foo", "Тема подписки")
	flag.StringVar(&User, "u", "postgres", "Имя пользователя СУБД")
	flag.StringVar(&Password, "p", "123", "Пароль пользователя СУБД")
	flag.StringVar(&DbName, "db", "TestL", "Имя подключаемой БД")
	flag.StringVar(&HTTPPort, "prt", "8080", "Порт работы http-сервера")

	log.SetFlags(0)
	flag.Parse()
	fmt.Printf("%s\n", usageStr)

	service := functional.Service{}
	go service.Start(&ClusterID, &ClientID, &Theme, &User, &Password, &DbName, &HTTPPort)
	fmt.Scanln()
}
