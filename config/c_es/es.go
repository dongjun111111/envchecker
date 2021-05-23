package c_es

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"goroot/config"
	"goroot/util"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v6"
)

type Obj_ES struct {
	ObjName string
}

func (s *Obj_ES) OutPut(v []byte, arg ...error) (res []byte) {
	s.ObjName = "[Elasticsearch]"
	return util.OutPut(s.ObjName, v, arg...)
}

func (s *Obj_ES) CheckObj(objcfg *config.ObjCfg) (res []byte) {
	defer objcfg.Wg.Done()
	db, err := NewEsDB(&Config{EsAddrs: []string{objcfg.Link}})
	if err != nil {
		return s.OutPut([]byte(objcfg.Link), errors.New("empty es link"))
	}
	body := map[string]interface{}{
		"test_idx": "test_idx",
	}
	indexid := fmt.Sprintf("%d.%d", "test_idx", "test_idx")
	jsonBody, _ := json.Marshal(body)
	if err := db.EsDbIndexWrite("test_idx", indexid, bytes.NewReader(jsonBody)); err != nil {
		log.Println("err:", err)
		return s.OutPut([]byte(objcfg.Link), errors.New("EsDbIndexWrite err :"+err.Error()))
	}
	return s.OutPut([]byte(objcfg.Link), nil)
}

type Config struct {
	EsAddrs []string
}

type ElDb struct {
	client *elasticsearch.Client
}

func NewEsDB(c *Config) (db *ElDb, err error) {
	cfg := elasticsearch.Config{
		Addresses: c.EsAddrs,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: util.DialTimeOutDuration,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
		return nil, err
	}

	db = &ElDb{
		client: es,
	}
	return db, nil
}

func (eldb *ElDb) EsDbIndexWrite(index, documentid string, msg io.Reader) (err error) {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: documentid,
		Body:       msg,
		Refresh:    "true",
		Timeout:    util.DialTimeOutDuration,
	}
	ctx, cancel := context.WithTimeout(context.Background(), util.DialTimeOutDuration)
	defer cancel()
	res, err := req.Do(ctx, eldb.client)
	if err != nil {
		log.Println("Error getting response: %s", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Println("[%s] Error indexing document ID=%s", res.Status(), documentid)

	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Println("Error parsing the response body: %s", err)
		}
	}
	return nil
}
