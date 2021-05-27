package c_hbase

import (
	"context"
	"goroot/config"
	"goroot/util"
	"log"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

type Obj_Hbase struct {
	ObjName string
}

func (s *Obj_Hbase) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Hbase]"
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Hbase) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	if objcfg.Link == "" {
		log.Println("empty hbase dsn")
		return
	}
	client := gohbase.NewClient(objcfg.Link)
	ctx, cancel := context.WithTimeout(context.Background(), util.DialTimeOutDuration)
	defer cancel()
	getRequest, err := hrpc.NewGetStr(ctx, "emp", "1")
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	_, err = client.Get(getRequest)
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	if client != nil {
		defer client.Close()
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
