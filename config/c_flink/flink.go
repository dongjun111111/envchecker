package c_flink

import (
	"goroot/config"
	"goroot/util"
	"log"

	"github.com/flink-go/api"
)

type Obj_Flink struct {
	ObjName string
}

func (s *Obj_Flink) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Flink]"
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Flink) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	if objcfg.Link == "" {
		log.Println("empty flink dsn")
		return
	}
	c, err := api.New(objcfg.Link)
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	_, err = c.Config()
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
