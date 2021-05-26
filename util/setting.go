package util

import (
	"log"
	"strings"

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

var config = struct {
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

var Config = config

func Setup(src string) error {
	var err error
	Config = config
	if src == "" {
		err = gcfg.ReadFileInto(&Config, "config/config.ini")
	} else {
		err = gcfg.ReadStringInto(&Config, strings.ReplaceAll(src, "update-config", ""))
		if err != nil {
			log.Println("Failed to parse "+src+":", err)
		}
	}
	if err != nil {
		log.Println("Failed to parse config file:", err)
	}
	return err
}
