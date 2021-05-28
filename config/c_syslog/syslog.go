package c_syslog

import (
	"errors"
	"goroot/config"
	"goroot/util"
	"log"
	"strings"

	gsyslog "github.com/hashicorp/go-syslog"
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
	_, err := gsyslog.DialLogger(linkArr[0], linkArr[1], gsyslog.LOG_ERR, "", "test")
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
