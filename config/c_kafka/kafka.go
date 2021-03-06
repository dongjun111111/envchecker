package c_kafka

import (
	"log"
	"goroot/config"
	"goroot/util"
)

type Obj_Kafka struct {
	ObjName string
}

func (s *Obj_Kafka) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[KafkaProducer]"
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Kafka) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	if objcfg.Link == "" {
		log.Println("empty kafka link")
		return
	}
	err := SendMessage(objcfg.Kafka_Data, objcfg.Kafka_TopicName)
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	b = append(b, []byte(" Sent kafka message:")...)
	b = append(b, objcfg.Kafka_Data...)
	return s.OutPut(b, nil)
}
