package handler

import (
	"main/model"
	"main/service"

	"github.com/gin-gonic/gin"
)

func Login(r *gin.Context) {
	typ := r.Query("type")

	switch typ {
	case "MuxiPass":
		//TODO
	case "ccnu":
		//TODO
	default:
		NativeLogin(r)
	}
}

func NativeLogin(r *gin.Context) {
	var loginAuth model.Auth
	r.ShouldBindJSON(&loginAuth)
	token, err := service.Login(loginAuth)
	if err != nil {
		SendError(r, err, nil,
			model.ErrorSender()+err.Error(), 403)
		return
	}
	SendResponse(r,nil,)
}
