package main

import (
	"errors"
	"goroot/config"
	"goroot/config/c_apm"
	"goroot/config/c_apollo"
	"goroot/config/c_clickhouse"
	"goroot/config/c_consul"
	"goroot/config/c_es"
	"goroot/config/c_flink"
	"goroot/config/c_hbase"
	"goroot/config/c_kafka"
	"goroot/config/c_mysql"
	"goroot/config/c_postgres"
	"goroot/config/c_redis"
	"goroot/config/c_syslog"
	"goroot/config/c_tidb"
	"goroot/util"
	"strconv"
	"sync"
	"time"

	melody "gopkg.in/olahol/melody.v1"
)

func ActionOnce(m *melody.Melody) {
	m.Broadcast([]byte("[BEGIN]"))
	m.Broadcast([]byte("Process:"))
	startTimestamp := time.Now()
	m.Broadcast([]byte("Start at:" + startTimestamp.Format("2006-01-02 15:04:05.000")))
	for {
		wg := &sync.WaitGroup{}
		var funcObj config.Check

		//mysql
		m_dsns := util.Config.Mysql.Link
		for i := 0; i < len(m_dsns); i++ {
			wg.Add(1)
			go func(i2 int, wgg *sync.WaitGroup) {
				funcObj = &c_mysql.Obj_Mysql{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: m_dsns[i2]})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(i, wg)
		}

		//redis
		r_dsns := util.Config.Redis.Link
		for i := 0; i < len(r_dsns); i++ {
			wg.Add(1)
			go func(i2 int, wgg *sync.WaitGroup) {
				funcObj = &c_redis.Obj_Redis{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: r_dsns[i2]})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(i, wg)
		}

		//es
		es_dsns := util.Config.ES.Link
		for i := 0; i < len(es_dsns); i++ {
			wg.Add(1)
			go func(i2 int, wgg *sync.WaitGroup) {
				funcObj = &c_es.Obj_ES{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: es_dsns[i2]})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(i, wg)
		}

		//consul
		cn_dsns := util.Config.Consul.Link
		for i := 0; i < len(cn_dsns); i++ {
			wg.Add(1)
			go func(i2 int, wgg *sync.WaitGroup) {
				funcObj = &c_consul.Obj_Consul{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: cn_dsns[i2]})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(i, wg)
		}

		//syslog
		sys_dsns := util.Config.Syslog.Link
		for i := 0; i < len(sys_dsns); i++ {
			wg.Add(1)
			go func(i2 int, wgg *sync.WaitGroup) {
				funcObj = &c_syslog.Obj_Syslog{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: sys_dsns[i2]})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(i, wg)
		}

		//clickhouse
		ch_dsns := util.Config.Clickhouse.Link
		for i := 0; i < len(ch_dsns); i++ {
			wg.Add(1)
			go func(i2 int, wgg *sync.WaitGroup) {
				funcObj = &c_clickhouse.Obj_Clickhouse{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: ch_dsns[i2]})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(i, wg)
		}

		//postgres
		pq_dsns := util.Config.Postgres.Link
		for i := 0; i < len(pq_dsns); i++ {
			wg.Add(1)
			go func(i2 int, wgg *sync.WaitGroup) {
				funcObj = &c_postgres.Obj_Postgres{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: pq_dsns[i2]})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(i, wg)
		}

		//flink
		fk_dsns := util.Config.Flink.Link
		for i := 0; i < len(fk_dsns); i++ {
			wg.Add(1)
			go func(i2 int, wgg *sync.WaitGroup) {
				funcObj = &c_flink.Obj_Flink{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: fk_dsns[i2]})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(i, wg)
		}

		//hbase
		hb_dsns := util.Config.Hbase.Link
		for i := 0; i < len(hb_dsns); i++ {
			wg.Add(1)
			go func(i2 int, wgg *sync.WaitGroup) {
				funcObj = &c_hbase.Obj_Hbase{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: hb_dsns[i2]})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(i, wg)
		}

		//tidb
		ti_dsns := util.Config.Tidb.Link
		for i := 0; i < len(ti_dsns); i++ {
			wg.Add(1)
			go func(i2 int, wgg *sync.WaitGroup) {
				funcObj = &c_tidb.Obj_Tidb{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: ti_dsns[i2]})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(i, wg)
		}

		//apm
		apm_dsns := util.Config.Apm.Link
		for i := 0; i < len(apm_dsns); i++ {
			wg.Add(1)
			go func(i2 int, wgg *sync.WaitGroup) {
				funcObj = &c_apm.Obj_Apm{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: apm_dsns[i2]})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(i, wg)
		}

		//apollo
		wg.Add(1)
		go func(wgg *sync.WaitGroup) {
			funcObj = &c_apollo.Obj_Apollo{}
			res := funcObj.CheckObj(&config.ObjCfg{
				Apollo_Link:          util.Config.Apollo.Link,
				Apollo_Appid:         util.Config.Apollo.Apollo_AppId,
				Apollo_Cluster:       util.Config.Apollo.Apollo_Cluster,
				Apollo_NamespaceName: util.Config.Apollo.Apollo_NamespaceName})
			if res != nil {
				m.Broadcast(res)
			}
			wgg.Done()
		}(wg)

		//kafka
		broker := util.Config.Kafka.Link
		kafkaConsumerCh := make(chan int)
		if broker != "" {
			topic := util.Config.Kafka.Kafka_Topic
			m.Broadcast(c_kafka.InitKafka(broker))
			go c_kafka.NewAccessLogConsumer(broker, topic, "group-1", m, kafkaConsumerCh)
			wg.Add(1)
			go func(wgg *sync.WaitGroup) {
				funcObj = &c_kafka.Obj_Kafka{}
				res := funcObj.CheckObj(&config.ObjCfg{
					Link: broker, Kafka_TopicName: topic,
					Kafka_Data: []byte("test-kafka-msg @ " + time.Now().Format("2006-01-02 15:04:05.000"))})
				if res != nil {
					m.Broadcast(res)
				}
				wgg.Done()
			}(wg)
		}
		wg.Wait()
		if broker != "" {
			select {
			case <-kafkaConsumerCh:
			case <-time.After(util.KafkaConsumerWaitDuration):
				m.Broadcast(util.OutPut("[KafkaConsumer]", []byte(broker), errors.New("Had waited 30 secs for kakfa-consumer,auto skip!")))
			}
		}
		goto END
	}
END:
	duration := time.Now().Sub(startTimestamp)
	m.Broadcast([]byte("Finish at:" + time.Now().Format("2006-01-02 15:04:05.000")))
	m.Broadcast([]byte(""))
	m.Broadcast([]byte(""))
	m.Broadcast([]byte("Summary:"))
	m.Broadcast([]byte("Total job number:" + strconv.Itoa(util.TotalFailedJob_Num+util.TotalSucceedJob_Num)))
	m.Broadcast([]byte("Total failed job number:" + strconv.Itoa(util.TotalFailedJob_Num)))
	m.Broadcast([]byte("Total succeed job number:" + strconv.Itoa(util.TotalSucceedJob_Num)))
	m.Broadcast([]byte("Total duration:" + duration.String()))
	m.Broadcast([]byte(""))
	if util.TotalFailedJob_Num > 0 {
		m.Broadcast([]byte("Total failed job detail:"))
		for i := 0; i < len(util.FailedProces); i++ {
			m.Broadcast(append([]byte("      "+strconv.Itoa(i+1)+"、"), util.FailedProces[i]...))
		}
	} else {
		m.Broadcast([]byte("Congratulations, all jobs are completed!"))
	}
	m.Broadcast([]byte("[END]"))
	util.InitNumbers()
}

func ActionAutoRefresh(m *melody.Melody) {
	for {
		select {
		case <-util.StopRefreshCh:
			m.Broadcast([]byte("auto-refresh has been suspended!"))
			goto END
		default:
		}
		m.Broadcast([]byte("auto-refresh"))
		ActionOnce(m)
		time.Sleep(util.AutoRefreshWaitDuration)
	}
END:
}
