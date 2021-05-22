package util

import (
	"bytes"
	"sync"
	"time"
)

var (
	TotalJob_Num        int
	TotalFailedJob_Num  int
	TotalSucceedJob_Num int
	FailedProces        [][]byte
	FLock               sync.Mutex
	SLock               sync.Mutex
	StopRefreshCh       chan int
	AutoRefreshStarted  bool
	AutoRefreshLock     sync.Mutex

	KafkaConsumerWaitDuration = 30 * time.Second
	AutoRefreshWaitDuration   = 5 * time.Second
	ApolloTouchWaitDuration   = 3 * time.Second

	Port = ":8080"
)

func init() {
	StopRefreshCh = make(chan int, 0)
}

func InitNumbers() {
	FailedProces = nil
	TotalJob_Num = 0
	TotalFailedJob_Num = 0
	TotalSucceedJob_Num = 0
}

func OutPut(objName string, v []byte, arg ...error) (res []byte) {
	if len(arg) > 0 && arg[0] != nil {
		FLock.Lock()
		defer FLock.Unlock()
		TotalFailedJob_Num++
		failed := []byte(objName + string(v) + arg[0].Error())
		FailedProces = append(FailedProces, failed)
		return failed
	}
	SLock.Lock()
	defer SLock.Unlock()
	TotalSucceedJob_Num++
	var resF []byte
	v = bytes.TrimPrefix(bytes.TrimSuffix(v, []byte(`"`)), []byte(`"`))
	v = append(v, []byte(" ->  Connected/Action succeed !")...)
	resF = []byte(objName)
	res = append(resF, v...)
	return
}
