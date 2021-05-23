package util

import (
	"log"

	"gopkg.in/gcfg.v1"
)

type Apm struct {
	Link []string
}

type Consul struct {
	Link []string
}

type ES struct {
	Link []string
}

type Mysql struct {
	Link []string
}

type Redis struct {
	Link []string
}

type Syslog struct {
	Link []string
}

type Kafka struct {
	Link        string
	Kafka_Topic string `gcfg:"Topic"`
}

type Apollo struct {
	Link                 string `gcfg:"Link"`
	Apollo_AppId         string `gcfg:"AppId"`
	Apollo_Cluster       string `gcfg:"Cluster"`
	Apollo_NamespaceName string `gcfg:"NamespaceName"`
}

type Clickhouse struct {
	Link []string
}
type Postgres struct {
	Link []string
}

var Config = struct {
	Apm        Apm
	Consul     Consul
	ES         ES
	Mysql      Mysql
	Redis      Redis
	Syslog     Syslog
	Kafka      Kafka
	Apollo     Apollo
	Clickhouse Clickhouse
	Postgres   Postgres
}{}

func Setup() {
	var err error
	err = gcfg.ReadFileInto(&Config, "config/config.ini")
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
}
