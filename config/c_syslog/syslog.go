package c_syslog

import (
	"errors"
	"log/syslog"
	"goroot/config"
	"goroot/util"
	"strings"
)

type Obj_Syslog struct {
	ObjName string
}

func (s *Obj_Syslog) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Syslog] "
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Syslog) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	defer objcfg.Wg.Done()
	if objcfg.Link == "" {
		return s.OutPut(nil, errors.New(objcfg.Link+" empty syslog dsn"))
	}
	linkArr := strings.Split(objcfg.Link, ":")
	if len(linkArr) != 2 {
		return s.OutPut(nil, errors.New(objcfg.Link+" wrong syslog dsn"))
	}
	_, err := syslog.Dial(linkArr[0], linkArr[1], syslog.LOG_ERR, "Saturday")
	if err != nil {
		return s.OutPut(nil, errors.New(objcfg.Link+" consul connect failed."+err.Error()))
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
