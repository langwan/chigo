package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type GreetRequest struct {
	Name string `json:"name"`
}

type GreetResponse struct {
	Message string `json:"message"`
}

func (api Api) Greet(c *gin.Context, request *GreetRequest) (*GreetResponse, error) {
	message := fmt.Sprintf("hi %s", request.Name)
	return &GreetResponse{Message: message}, nil
}
