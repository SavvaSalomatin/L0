package main

import (
	"database/sql"
	"github.com/nats-io/stan.go"
	"servicel0/config"
	"servicel0/functional"
	"testing"
)

func TestValidator(t *testing.T) {
	bad := false
	good := true
	var unvalidByteStream [5][]byte
	var validByteStream [5][]byte
	unvalidByteStream[0] = []byte("sldkgslkgha;sldghas;lgh")
	unvalidByteStream[1] = []byte("{\n\t\t\"order_uid\": \"b21\",\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 1500,\n\t\t  \"goods_total\": 317,\n\t\t  \"custom_fee\": 0\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": \"28\",\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
	unvalidByteStream[2] = []byte("{\n\t\t\"order_uid\": \"b21\",\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 1500,\n\t\t  \"goods_total\": 317,\n\t\t  \"custom_fee\": 10\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\"\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
	unvalidByteStream[3] = []byte("{\n\t\t\"order_uid\": \"b21dfgdfgbnfgnfgnghmfjghlhjv,mjh,hjv,\",\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 1500,\n\t\t  \"goods_total\": 317,\n\t\t  \"custom_fee\": \"sdf\"\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_crated\": \"2021-11-2606:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
	unvalidByteStream[4] = []byte("{\n\t\t\"order_uid\": \"\",\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 1500,\n\t\t  \"goods_total\": 317,\n\t\t  \"custom_fee\": \"sdf\"\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
	validByteStream[0] = []byte("{\n\t\t\"order_uid\": \"b21dfhdf435\",\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 1500,\n\t\t  \"goods_total\": 317,\n\t\t  \"custom_fee\": 0\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
	validByteStream[1] = []byte("{\n\t\t\"order_uid\": \"b21dfhdf435dfbcv\",\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 1500,\n\t\t  \"goods_total\": 317,\n\t\t  \"custom_fee\": 10\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
	validByteStream[2] = []byte("{\n\t\t\"order_uid\": \"b21dfhdf435sdf\",\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 1500,\n\t\t  \"goods_total\": 317,\n\t\t  \"custom_fee\": 0\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
	validByteStream[3] = []byte("{\n\t\t\"order_uid\": \"b21dfhdf435sdfdfhg\",\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 13489573945,\n\t\t  \"goods_total\": 366,\n\t\t  \"custom_fee\": 28\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
	validByteStream[4] = []byte("{\n\t\t\"order_uid\": \"b\",\n\t\t\"track_number\": \"WBILMfghfghTESTTRACK\",\n\t\t\"entry\": \"WBIL\",\n\t\t\"delivery\": {\n\t\t  \"name\": \"Test Testov\",\n\t\t  \"phone\": \"+9720000000\",\n\t\t  \"zip\": \"2639809\",\n\t\t  \"city\": \"Kiryat Mozkin HHH\",\n\t\t  \"address\": \"Ploshad Mira 15\",\n\t\t  \"region\": \"Kraidfgot\",\n\t\t  \"email\": \"test@gmail.com\"\n\t\t},\n\t\t\"payment\": {\n\t\t  \"transaction\": \"b563uyiiu7bb84b\",\n\t\t  \"request_id\": \"\",\n\t\t  \"currency\": \"USD\",\n\t\t  \"provider\": \"wbpay\",\n\t\t  \"amount\": 1817,\n\t\t  \"payment_dt\": 1637907727,\n\t\t  \"bank\": \"alpha\",\n\t\t  \"delivery_cost\": 0,\n\t\t  \"goods_total\": 340,\n\t\t  \"custom_fee\": 15\n\t\t},\n\t\t\"items\": [\n\t\t  {\n\t\t\t\"chrt_id\": 9934930,\n\t\t\t\"track_number\": \"WBILMTESTTRACK\",\n\t\t\t\"price\": 453,\n\t\t\t\"rid\": \"ab4219087a764ae0btest\",\n\t\t\t\"name\": \"Mascaras\",\n\t\t\t\"sale\": 30,\n\t\t\t\"size\": \"0\",\n\t\t\t\"total_price\": 317,\n\t\t\t\"nm_id\": 2389212,\n\t\t\t\"brand\": \"Vivienne Sabo\",\n\t\t\t\"status\": 202\n\t\t  }\n\t\t],\n\t\t\"locale\": \"en\",\n\t\t\"internal_signature\": \"\",\n\t\t\"customer_id\": \"test\",\n\t\t\"delivery_service\": \"meest\",\n\t\t\"shardkey\": \"9\",\n\t\t\"sm_id\": 99,\n\t\t\"date_created\": \"2021-11-26T06:22:19Z\",\n\t\t\"oof_shard\": \"1\"\n\t  }")
	for i := 0; i < 5; i++ {
		if bad != functional.Validcheck(unvalidByteStream[i]) {
			t.Errorf("Требуется: %t;", bad)
		}
	}
	for i := 0; i < 5; i++ {
		if good != functional.Validcheck(validByteStream[i]) {
			t.Errorf("Требуется: %t;", good)
		}
	}
}

func TestConnectdb(t *testing.T) {
	var user = "postgres"
	var pass = "123"
	var dbn = "TestL"
	pointerdb := functional.Connectdb(&user, &pass, &dbn)
	err := pointerdb.Ping()
	if err != nil {
		t.Errorf("Требуется корректный указатель, ошибка: %p;", err)
	}
}

func TestDocache(t *testing.T) {
	var confstrpostgres = "user=postgres password=123 dbname=TestL sslmode=disable"
	db, _ := sql.Open("postgres", confstrpostgres)
	err := functional.Docache(db)
	if config.C == nil {
		t.Errorf("Кеш пуст, ошибка: %p;", err)
	}
}

func TestStanConnect(t *testing.T) {
	sc, aw := functional.StanConnect("test-cluster", "clientidid")
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		m.Ack()
	}, stan.SetManualAckMode(), stan.AckWait(aw), stan.DurableName("123"))
	sub.Unsubscribe()
	sc.Close()
	if err != nil {
		t.Errorf("Не удалось подключиться к серверу stan: %p;", err)
	}
}
