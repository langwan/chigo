package main

import "github.com/gin-gonic/gin"

func main() {
	g := gin.New()
	g.Use(Validator())
	g.POST("/login", Login)
	g.Run(":8080")
}
