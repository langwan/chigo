package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
)

type Data struct {
	Name string
}

const (
	HostSuffix = ".faming.com:9100"
)

func main() {
	g := gin.New()
	g.Any("", func(c *gin.Context) {
		hostname := ""
		//http://chihuo.faming.com:9100/
		if strings.HasSuffix(c.Request.Host, HostSuffix) {
			index := strings.LastIndex(c.Request.Host, HostSuffix)
			hostname = c.Request.Host[0:index]
		}
		if len(hostname) > 0 {
			data := Data{Name: hostname}
			tmpl, err := template.ParseFiles("./html/profile.html")
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(c.Writer, data)
			if err != nil {
				panic(err)
			}
		} else {
			tmpl, err := template.ParseFiles("./html/home.html")
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(c.Writer, nil)
			if err != nil {
				panic(err)
			}
		}

	})
	g.Run(":9100")
}
