package handler

import (
	"main/model"
	"main/model/db"
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
	case "qq":
		//TODO
	default:
		NativeLogin(r)
	}
}
// NativeLogin godoc
//	@Summary		native login
//	@Description	login and return id&token
//	@Param			loginAuth	formData	object	true	"the User who is logining"
//	@Produce		json
//	@Success		200	{object}	handler.Response	"id&token"
//	@Failure		404	{object}	handler.Response
//	@Router			/login [post]
func NativeLogin(r *gin.Context) {
	var loginAuth db.User
	r.ShouldBindJSON(&loginAuth)
	token, err := service.LoginNative(loginAuth)
	if err != nil {
		SendError(r, err, nil,
			model.ErrorSender()+err.Error(), 403)
		return
	}
	SendResponse(r, nil, gin.H{
		"id":    loginAuth.UID,
		"token": token,
	})
}
