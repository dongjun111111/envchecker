package c_apm

import (
	"errors"
	"goroot/config"
	"goroot/util"
	"log"
	"net/url"

	"go.elastic.co/apm"
	"go.elastic.co/apm/transport"
)

type Obj_Apm struct {
	ObjName string
}

func (s *Obj_Apm) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Apm] "
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Apm) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	defer objcfg.Wg.Done()
	if objcfg.Link == "" {
		return s.OutPut([]byte(objcfg.Link), errors.New("empty apm link"))
	}
	trans, err := transport.NewHTTPTransport()
	if err != nil {
		log.Println("tracer create transport failed.", err)
		return s.OutPut([]byte(objcfg.Link), errors.New("tracer create transport failed."+err.Error()))
	}
	u, err := url.Parse(objcfg.Link)
	if err != nil {
		log.Println("get apm-server error.", err)
		return s.OutPut([]byte(objcfg.Link), errors.New("get apm-server error."+err.Error()))
	}
	trans.SetServerURL(u)
	apm.DefaultTracer.Transport = trans
	apm.DefaultTracer.SetRequestDuration(util.DialTimeOutDuration)
	transaction := apm.DefaultTracer.StartTransaction("GET /", "request")
	if transaction.Result != "Success" {
		log.Println("apm-server response not ok.")
		return s.OutPut([]byte(objcfg.Link), errors.New("apm-server response not ok."+transaction.Result))
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
