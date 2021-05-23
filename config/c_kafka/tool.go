package c_kafka

import (
	"encoding/binary"
	"errors"
	"goroot/util"
	"log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"gopkg.in/olahol/melody.v1"
)

func Int64ToBytes(uid int64) []byte {
	var b []byte
	binary.BigEndian.PutUint64(b, uint64(uid))
	return b
}

var businessProducerObj sarama.AsyncProducer
var err error

func InitKafka(broker string) []byte {
	businessProducerObj, err = newAccessLogProducer([]string{broker})
	if err != nil {
		return util.OutPut("[KafkaInit]", []byte(broker), err)
	}
	return util.OutPut("[KafkaInit]", []byte(broker+" init kafka succeed!"), nil)
}
func newAccessLogProducer(brokerList []string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Timeout = util.DialTimeOutDuration
	config.Net.DialTimeout = util.DialTimeOutDuration
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		log.Println("start-kafka-client-failed.", err)
		return nil, err
	}
	go func() {
		for err := range producer.Errors() {
			log.Println("kafka-producer-init-failed.", err)
		}
	}()
	return producer, nil
}

type accessLogEntry struct {
	encoded []byte
}

func (ale *accessLogEntry) Length() int {
	return len(ale.encoded)
}

func (ale *accessLogEntry) Encode() ([]byte, error) {
	return ale.encoded, nil
}

func SendMessage(data []byte, topicName string) error {
	if businessProducerObj == nil {
		return errors.New("businessProducerObj == nil")
	}
	businessProducerObj.Input() <- &sarama.ProducerMessage{
		Topic: topicName,
		Value: &accessLogEntry{encoded: data},
	}
	return nil
}

func NewAccessLogConsumer(broker string, topics string, groupId string, m *melody.Melody, kafkaConsumerCh chan int) {
	if businessProducerObj != nil {
		defer businessProducerObj.Close()
	}
	config := cluster.NewConfig()
	config.Config.Net.DialTimeout = util.DialTimeOutDuration
	config.Consumer.MaxWaitTime = util.DialTimeOutDuration
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.CommitInterval = time.Second

	topicslice := strings.Split(topics, ",")
	consumer, err := cluster.NewConsumer([]string{broker}, groupId, topicslice, config)
	if err != nil {
		log.Printf("%s: sarama.NewSyncProducer err, message=%s \n", groupId, err)
		var b []byte
		b = append(b, []byte(broker)...)
		b = util.OutPut("[KafkaConsumer]", b, err)
		m.Broadcast(b)
		kafkaConsumerCh <- 1
		return
	}
	defer consumer.Close()

	go func() {
		for err := range consumer.Errors() {
			log.Printf("%s:Error: %s\n", groupId, err.Error())
		}
	}()

	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("%s:Rebalanced: %+v \n", groupId, ntf)
		}
	}()

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				consumer.MarkOffset(msg, "")
				log.Println("rcv kafka messages : ", msg.Topic, string(msg.Key), string(msg.Value))
				var b []byte
				b = append(b, []byte(broker)...)
				b = append(b, []byte(" Received kafka message:")...)
				b = append(b, msg.Value...)
				b = util.OutPut("[KafkaConsumer]", b, nil)
				m.Broadcast(b)
				kafkaConsumerCh <- 1
				goto END
			}
		case <-time.After(util.KafkaConsumerWaitDuration):
			var b []byte
			b = append(b, []byte(broker)...)
			b = util.OutPut("[KafkaConsumer]", b, errors.New("30 secs timeout ,auto exit"))
			m.Broadcast(b)
			goto END
		}
	}
END:
}
