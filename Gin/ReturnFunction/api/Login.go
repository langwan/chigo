package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Name string `json:"name"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

func (api Api) Login(c *gin.Context, request *LoginRequest) (*LoginResponse, error) {
	message := fmt.Sprintf("hi %s, login ok", request.Name)
	return &LoginResponse{Message: message}, nil
}
