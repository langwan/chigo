package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/langwan/chigo/Gin/ReturnFunction/api"
	helper_code "github.com/langwan/langgo/helpers/code"
	"io"
)

func main() {
	g := gin.Default()
	g.Any("api/*uri", ApiHandler())
	g.Any("hello", HelloHandler)
	g.Run(":9010")
}

func ApiHandler() gin.HandlerFunc {
	a := api.Api{}
	return func(c *gin.Context) {
		methodName := c.Param("uri")[1:]

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			return
		}

		if err != nil {
			c.AbortWithStatus(500)
			return
		}

		response, code, err := helper_code.Call(c, &a, methodName, string(body))

		if err != nil {
			c.AbortWithError(500, err)
		} else if code != 0 {
			a.SendBad(c, code, err.Error(), nil)
		} else {
			a.SendOk(c, response)
		}
	}
}

func HelloHandler(c *gin.Context) {
	if true {
		c.AbortWithStatus(403)
		return
	}
	fmt.Println("end")
}
