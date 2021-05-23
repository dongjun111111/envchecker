package c_postgres

import (
	"database/sql"
	"errors"
	"goroot/config"
	"goroot/util"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Obj_Postgres struct {
	ObjName string
}

func (s *Obj_Postgres) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[PostgreSQL]"
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Postgres) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	defer objcfg.Wg.Done()
	if objcfg.Link == "" {
		return s.OutPut([]byte(objcfg.Link), errors.New("empty PostgreSQL dsn"))
	}
	if !strings.Contains(objcfg.Link, "timeout") {
		objcfg.Link += "&timeout=3s"
	}
	sqlDB, err := sql.Open("postgres", objcfg.Link)
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	if sqlDB != nil {
		defer sqlDB.Close()
	}
	_, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
