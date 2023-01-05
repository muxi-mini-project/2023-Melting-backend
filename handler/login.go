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
	s := loginAuth.Auth
	model.GetSth(&loginAuth)
	if s == loginAuth.Auth {
		t, err := service.Newtoken(int(loginAuth.UID))
		if err != nil {
			r.JSON(500, err)
		} else {
			r.JSON(200, gin.H{
				"code":  200,
				"token": t,
			})
		}
	} else {
		r.String(403, "Wrong username or password")
	}
}
