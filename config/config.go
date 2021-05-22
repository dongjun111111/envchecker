package config

import "sync"

type ObjCfg struct {
	Wg   *sync.WaitGroup
	Link string

	Kafka_TopicName string
	Kafka_Data      []byte

	Apollo_Appid         string
	Apollo_Link          string
	Apollo_Cluster       string
	Apollo_NamespaceName string
}

type Check interface {
	CheckObj(objcfg *ObjCfg) (res []byte)
	OutPut(v []byte, arg ...error) (res []byte)
}
