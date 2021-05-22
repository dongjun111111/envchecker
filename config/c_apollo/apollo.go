package c_apollo

import (
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
	s.ObjName = "[Apollo] "
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Apollo) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	defer objcfg.Wg.Done()
	url := objcfg.Apollo_Link + "/configs/" + objcfg.Apollo_Appid + "/" + objcfg.Apollo_Cluster + "/" + objcfg.Apollo_NamespaceName
	if !strings.Contains(objcfg.Apollo_Link, "http") {
		url = "http://" + url
	}
	c := &http.Client{
		Timeout: util.ApolloTouchWaitDuration,
	}
	resp, err := c.Get(url)
	if err != nil || resp == nil || (resp != nil && resp.StatusCode != 200) {
		if err != nil {
			return s.OutPut([]byte(url), errors.New(" apollo failed."+err.Error()))
		}
		return s.OutPut([]byte(url), errors.New("  apollo failed."+resp.Status))
	}
	return s.OutPut([]byte(url), nil)
}
