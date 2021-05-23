package c_mysql

import (
	"errors"
	"goroot/config"
	"goroot/util"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Obj_Mysql struct {
	ObjName string
}

func (s *Obj_Mysql) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[MySQL]"
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Mysql) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	defer objcfg.Wg.Done()
	if objcfg.Link == "" {
		return s.OutPut([]byte(objcfg.Link), errors.New("empty mysql dsn"))
	}
	if !strings.Contains(objcfg.Link, "timeout") {
		objcfg.Link += "&timeout=3s"
	}
	_, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       objcfg.Link,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
