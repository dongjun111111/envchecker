package c_tidb

import (
	"goroot/config"
	"goroot/util"
	"log"

	"github.com/pingcap/tidb/store/tikv"
)

type Obj_Tidb struct {
	ObjName string
}

func (s *Obj_Tidb) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Tidb]"
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Tidb) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	if objcfg.Link == "" {
		log.Println("empty tidb dsn")
		return
	}
	driver := tikv.Driver{}
	_, err := driver.Open(objcfg.Link)
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
