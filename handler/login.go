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
//	@Param			loginAuth	body	model.LoginRequest	true	"the User who is logging in"
//	@Produce		json
//	@Success		200	{object}	model.LoginResponse
//	@Failure		401	{object}	handler.Response	"username or password incorrect"
//	@Failure		403	{object}	handler.Response	"param not satisfied"
//	@Failure		500	{object}	handler.Response	"token generation failed"
//	@Router			/login [post]
func NativeLogin(r *gin.Context) {
	loginAuth := db.User{}
	err := r.ShouldBindJSON(&loginAuth)
	if err != nil {
		SendError(r, model.ErrBadRequest, loginAuth,
			model.ErrorSender(), http.StatusBadRequest)
		return
	}
	if loginAuth.Auth == "" || loginAuth.NickName == "" {
		SendError(r, model.ErrAuthInvalid, loginAuth,
			model.ErrorSender(), http.StatusBadRequest)
		return
	}
	token, err := service.LoginNative(loginAuth)
	if err == model.ErrAuthIncorrect {
		SendError(r, err, nil,
			model.ErrorSender(), http.StatusUnauthorized)
		return
	}
	if err != nil {
		SendError(r, err, nil,
			model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, model.LoginResponse{
		Code:  http.StatusAccepted,
		ID:    int(loginAuth.UID),
		Token: token,
	})
}
