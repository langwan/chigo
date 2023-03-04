package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 绑定 JSON
type Login struct {
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required,len=32"`
}

func main() {
	router := gin.Default()

	// 绑定 JSON ({"user": "manu", "password": "123"})
	router.POST("/login", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.Run(":8080")
}
