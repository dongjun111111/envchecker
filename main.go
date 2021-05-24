package main

import (
	"goroot/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	melody "gopkg.in/olahol/melody.v1"
)

func main() {
	util.Setup("")
	r := gin.Default()
	m := melody.New()
	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "template/index.html")
	})
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		if strings.Contains(string(msg), "update-config") {
			str := strings.ReplaceAll(string(msg), "update-config", "")
			err := util.Setup(str)
			if err != nil {
				m.Broadcast([]byte("update-config failed," + err.Error()))
			} else {
				m.Broadcast([]byte("update-config succeed!"))
			}
			return
		}
		switch string(msg) {
		case "start":
			ActionOnce(m)
		case "auto-refresh":
			util.AutoRefreshLock.Lock()
			defer util.AutoRefreshLock.Unlock()
			if util.AutoRefreshStarted {
				util.AutoRefreshStarted = false
				util.StopRefreshCh <- 1
				m.Broadcast([]byte("auto-refresh had been killed right now,please try again!"))
				return
			}
			util.AutoRefreshStarted = true
			go ActionAutoRefresh(m)
		case "suspend":
			util.AutoRefreshLock.Lock()
			defer util.AutoRefreshLock.Unlock()
			util.AutoRefreshStarted = false
			util.StopRefreshCh <- 1
		default:
		}
	})

	r.Run(util.Port)
}
