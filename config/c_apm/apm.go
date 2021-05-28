package c_apm

import (
	"log"
	"goroot/config"
	"goroot/util"
)

type Obj_Apm struct {
	ObjName string
}

func (s *Obj_Apm) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Apm]"
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Apm) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	if objcfg.Link == "" {
		log.Println("empty apm link")
		return
	}
	err := util.NetSniffer(objcfg.Link, "tcp")
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
