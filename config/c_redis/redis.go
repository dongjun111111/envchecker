package c_redis

import (
	"context"
	"errors"
	"goroot/config"
	"goroot/util"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type Obj_Redis struct {
	ObjName string
}

func (s *Obj_Redis) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Redis] "
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_Redis) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	defer objcfg.Wg.Done()
	if objcfg.Link == "" {
		return s.OutPut(nil, errors.New("empty redis dsn"))
	}

	u, err := url.Parse(objcfg.Link)
	if err != nil {
		return s.OutPut(nil, errors.New(objcfg.Link+" redis link parse failed."))
	}

	if u.Scheme == "redis" {
		opts := &redis.Options{
			Addr: u.Host,
			DB:   0,
		}
		if u.User != nil {
			opts.Password = u.User.Username()
		}
		dbstr := u.Query().Get("db")
		if dbstr != "" {
			db, err := strconv.Atoi(dbstr)
			if err != nil {
				return s.OutPut(nil, errors.New(objcfg.Link+" redis db parse error."+err.Error()))
			}
			opts.DB = db
		}
		rClient := redis.NewClient(opts)
		if rClient != nil {
			defer rClient.Close()
			ctx := context.Background()
			res := rClient.Ping(ctx)
			if res == nil || res.Err() != nil {
				return s.OutPut(nil, errors.New(objcfg.Link+" redis init client failed."))
			}
		} else {
			return s.OutPut(nil, errors.New(objcfg.Link+" redis-opts link error."))
		}
	} else if u.Scheme == "redis-cluster" {
		addrs := strings.Split(u.Host, ",")
		opts := &redis.ClusterOptions{
			Addrs:       addrs,
			Password:    u.User.Username(),
			PoolSize:    50,
			IdleTimeout: time.Minute * 30,
		}
		rClient := redis.NewClusterClient(opts)
		if rClient != nil {
			defer rClient.Close()
			ctx := context.Background()
			res := rClient.Ping(ctx)
			if res == nil || res.Err() != nil {
				return s.OutPut(nil, errors.New(objcfg.Link+" redis-cluster connect failed."))
			}
		} else {
			return s.OutPut(nil, errors.New(objcfg.Link+" redis-cluster-opts link error."))
		}
	} else {
		return s.OutPut(nil, errors.New(objcfg.Link+" invailed redis mode"))
	}
	var b []byte
	b = append(b, []byte(objcfg.Link)...)
	return s.OutPut(b, nil)
}
