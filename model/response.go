package model

import (
	"github.com/gin-gonic/gin"
)

var NoResponse = gin.H{
	"code":    200,
	"message": "null",
}

type LoginResponse struct {
	Code    int         
	ID      int
	Token   string
}
