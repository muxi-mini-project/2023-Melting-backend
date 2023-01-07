package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
} //@name Response

func SendResponse(c *gin.Context, err error, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "OK",
		Data:    data,
	})
}

func SendBadRequest(c *gin.Context, err error, data interface{}, cause string) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    403,
		Message: cause + err.Error(),
		Data:    data,
	})
}

func SendError(c *gin.Context, err error, data interface{}, cause string, code int) {
	c.JSON(code, Response{
		Code:    code,
		Message: cause + err.Error(),
		Data:    data,
	})
}
