package c_tidb

import (
	"context"
	"goroot/config"
	"goroot/util"
	"log"

	tiConfig "github.com/tikv/client-go/config"
	"github.com/tikv/client-go/rawkv"
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
	ctx, cancel := context.WithTimeout(context.Background(), util.DialTimeOutDuration)
	defer cancel()
	cli, err := rawkv.NewClient(ctx, []string{objcfg.Link}, tiConfig.Default())
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	defer cli.Close()
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
