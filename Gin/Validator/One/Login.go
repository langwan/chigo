package main

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Phone    string `json:"phone" binding:"required,len=11" label:"手机号"`
	Password string `json:"password" binding:"required,len=32" label:"密码"`
}

func Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		ValidatorErrors(c, err)
		return
	}
}
