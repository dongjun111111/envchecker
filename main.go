package main

import (
	"goroot/util"
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	melody "gopkg.in/olahol/melody.v1"
)

func initTemplates() *template.Template {
	box := packr.NewBox("./template")
	t := template.New("")
	tmpl := t.New("index.html")
	data, _ := box.FindString("index.html")
	tmpl.Parse(data)
	return t
}

func main() {
	util.Setup("")
	r := gin.Default()
	m := melody.New()

	r.SetHTMLTemplate(initTemplates())
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
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
