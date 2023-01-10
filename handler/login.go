package handler

import (
	"main/model"
	"main/model/db"
	"main/service"
	"net/http"

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
//
//	@Summary		native login
//	@Description	login and return id&token
//	@Accept			application/json
//	@Param			loginAuth	body	model.LoginRequest	true	"the User who is logining"
//	@Produce		json
//	@Success		200	{object}	model.LoginResponse
//	@Failure		404	{object}	handler.Response
//	@Router			/login [post]
func NativeLogin(r *gin.Context) {
	loginAuth := db.User{}
	r.ShouldBindJSON(&loginAuth)
	if loginAuth.Auth == "" || loginAuth.NickName == "" {
		SendError(r, model.ErrAuthInvalid, nil,
			model.ErrorSender(), http.StatusBadRequest)
		return
	}
	token, err := service.LoginNative(loginAuth)
	if err != nil {
		SendError(r, err, nil,
			model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, nil, model.LoginResponse{
		Code:  http.StatusAccepted,
		ID:    int(loginAuth.UID),
		Token: token,
	})
}
