package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

type Home struct {
	Title string
	Name  string
}

func main() {
	var err error
	defer func() {
		if err != nil {
			panic(err)
		}
	}()

	g := gin.New()
	g.GET("home", func(c *gin.Context) {
		tpl, err := template.ParseFiles("./tpl/layout.html", "./tpl/header.html", "./tpl/footer.html", "./pages/home.html")
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		tpl.Execute(c.Writer, &Home{
			Title: "home",
			Name:  "chihuo",
		})

	})
	g.GET("about", func(c *gin.Context) {
		tpl, err := template.ParseFiles("./tpl/layout.html", "./tpl/header.html", "./tpl/footer.html", "./pages/about.html")
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		tpl.Execute(c.Writer, &Home{
			Title: "about",
			Name:  "chihuo",
		})

	})
	g.Run(":8100")
}
