package c_syslog

import (
	"errors"
        "log"
	"goroot/config"
	"goroot/util"
	"log/syslog"
	"strings"
)

type Obj_Syslog struct {
	ObjName string
}

func (s *Obj_Syslog) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Syslog]"
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Syslog) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	if objcfg.Link == "" {
		log.Println("empty syslog dsn")
		return
	}
	linkArr := strings.Split(objcfg.Link, ":")
	if len(linkArr) != 2 {
		return s.OutPut([]byte(objcfg.Link), errors.New("wrong syslog dsn"))
	}
	_, err := syslog.Dial(linkArr[0], linkArr[1], syslog.LOG_ERR, "Saturday")
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
