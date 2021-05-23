package c_consul

import (
	"errors"
	"goroot/config"
	"goroot/util"
	"log"

	"github.com/hashicorp/consul/api"
)

type Obj_Consul struct {
	ObjName string
}

func (s *Obj_Consul) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Consul]"
	return util.OutPut(s.ObjName, v, arg...)
}
func (s *Obj_Consul) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	defer objcfg.Wg.Done()
	if objcfg.Link == "" {
		return s.OutPut([]byte(objcfg.Link), errors.New("empty consul link"))
	}
	config := api.DefaultConfig()
	config.Address = objcfg.Link
	config.WaitTime = util.DialTimeOutDuration
	if config.HttpClient != nil {
		config.HttpClient.Timeout = util.DialTimeOutDuration
	}
	client, err := api.NewClient(config)
	if err != nil {
		log.Println("checkconsul initclient failed, err:", err)
		return s.OutPut([]byte(objcfg.Link), err)
	}
	pair := &api.KVPair{
		Key: "env",
	}
	kv := client.KV()
	if p, _, err := kv.Get(pair.Key, nil); err != nil {
		log.Println("checkconsul failed, err:", err)
		return s.OutPut([]byte(objcfg.Link), err)
	} else {
		log.Println(p.Key, "=>", string(p.Value))
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
