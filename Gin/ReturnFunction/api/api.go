package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Body    interface{} `json:"body" `
}

type Api struct {
}

func (api Api) SendOk(c *gin.Context, body any) {
	resp := ApiResponse{}
	resp.Code = 0
	resp.Message = ""
	resp.Body = body
	c.JSON(http.StatusOK, resp)
}

func (api Api) SendBad(c *gin.Context, code int, message string, body any) {
	resp := ApiResponse{}
	resp.Code = code
	resp.Message = message
	resp.Body = body
	c.AbortWithStatusJSON(http.StatusOK, resp)
}
