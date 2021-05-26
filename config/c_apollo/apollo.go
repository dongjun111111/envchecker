package c_apollo

import (
	"log"
	"errors"
	"goroot/config"
	"goroot/util"
	"net/http"
	"strings"
)

type Obj_Apollo struct {
	ObjName string
}

func (s *Obj_Apollo) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Apollo]"
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Apollo) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	if objcfg.Apollo_Link == "" {
		log.Println("empty apollo link")
		return
	}
	url := objcfg.Apollo_Link + "/configs/" + objcfg.Apollo_Appid + "/" + objcfg.Apollo_Cluster + "/" + objcfg.Apollo_NamespaceName
	if !strings.Contains(objcfg.Apollo_Link, "http") {
		url = "http://" + url
	}
	c := &http.Client{
		Timeout: util.DialTimeOutDuration,
	}
	resp, err := c.Get(url)
	if err != nil || resp == nil || (resp != nil && resp.StatusCode != 200) {
		if err != nil {
			return s.OutPut([]byte(url), err)
		}
		return s.OutPut([]byte(url), errors.New("connect apollo failed."+resp.Status))
	}
	return s.OutPut([]byte(url), nil)
}
