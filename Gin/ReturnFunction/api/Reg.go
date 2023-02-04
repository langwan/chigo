package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RegRequest struct {
	Name string `json:"name"`
}

type RegResponse struct {
	Message string `json:"message"`
}

func (api Api) Reg(c *gin.Context, request *RegRequest) (*RegResponse, error) {
	message := fmt.Sprintf("hi %s, reg ok", request.Name)
	return &RegResponse{Message: message}, nil
}
