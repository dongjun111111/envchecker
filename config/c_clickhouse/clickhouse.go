package c_clickhouse

import (
	"errors"
	"fmt"
	"goroot/config"
	"goroot/util"
	"log"

	"database/sql"

	"github.com/ClickHouse/clickhouse-go"
)

type Obj_Clickhouse struct {
	ObjName string
}

func (s *Obj_Clickhouse) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Clickhouse]"
	return util.OutPut(s.ObjName, v, arg...)
}
func (s *Obj_Clickhouse) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	if objcfg.Link == "" {
		log.Println("empty clickhouse link")
		return
	}
	connect, err := sql.Open("clickhouse", objcfg.Link)
	if err != nil {
		log.Println(err)
		return s.OutPut([]byte(objcfg.Link), err)
	}
	if connect != nil {
		defer connect.Close()
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			errStr := fmt.Sprintf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
			log.Println(exception.Code, exception.Message, exception.StackTrace)
			return s.OutPut([]byte(objcfg.Link), errors.New(errStr))
		} else {
			log.Println(err)
			return s.OutPut([]byte(objcfg.Link), err)
		}
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
