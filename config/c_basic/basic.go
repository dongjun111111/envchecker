package c_basic

import (
	"errors"
	"goroot/config"
	"goroot/util"
	"log"
	"net"
	"strings"
)

type Obj_Basic struct {
	ObjName string
}

func (s *Obj_Basic) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Basic]"
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Basic) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	if objcfg.Link == "" {
		log.Println("empty basic dsn")
		return
	}
	err := netSniffer(objcfg.Link, "tcp")
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), err)
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}

func netSniffer(addr, netType string) (err error) {
	if len(strings.Split(addr, ":")) != 2 {
		err = errors.New(" wrong addr!")
		return
	}
	netType = strings.ToLower(netType)
	if netType != "udp" && netType != "tcp" {
		err = errors.New(netType + " wrong net type!")
		return
	}
	conn, err := net.DialTimeout(netType, addr, util.DialTimeOutDuration)
	if err != nil {
		return
	}
	if conn == nil {
		err = errors.New(addr + " conn failed!")
		return
	}
	_ = conn.Close()
	return nil
}
