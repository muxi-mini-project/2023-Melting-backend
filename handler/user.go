package handler

import (
	"main/model"
	"main/model/db"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadProfile godoc
//	@Summary		upload profile
//	@Tags			user
//	@Description	upload sth with its UserID
//	@Accept			application/json
//	@Produce		json
//	@Param			Authorization	header		string	true	"token"
//	@Success		200		{string}	string
//	@Router			/users [put]
func UploadProfile(r *gin.Context) {
	id := r.GetInt("userID")
	data := db.User{
		UID: int32(id),
	}
	r.ShouldBindJSON(&data)
	model.UpdateSth(data)
	SendResponse(r, nil, model.NoResponse)
}

// UploadPhoto godoc
//	@Summary		upload photo
//	@Tags			user
//	@Description	upload photo with its UserID
//	@Accept			image/jpeg
//	@Produce		json
//	@Param			file	formData	object	true	"the photo of the user"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200		{object}	string
//	@Router			/users/photo [put]
func UploadPhoto(r *gin.Context) {
	id := r.GetInt("userID")
	H, err := r.FormFile("file")
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
	}
	file, err := H.Open() // Warning: file must be *.jpg
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusForbidden)
	}
	url, err := service.UploadProfilePhoto(id, file, H.Size)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusInternalServerError)
	}
	var data = db.User{
		UID: int32(id),
	}
	data = model.GetSth(data)
	data.Photo = url
	model.UpdateSth(data)
	SendResponse(r, nil, data)
}

// GetUserInfo godoc
//	@Summary		Get User's info
//	@Description	Get User's info with its userID
//	@Param			Authorization	header	string	true	"token"
//	@Produce		json
//	@Success		200	{object}	db.User
//	@Failure		404	{object}	handler.Response
//	@Router			/users [get]
func GetUserInfo(r *gin.Context) {
	id := r.GetInt("userID")
	data := db.User{
		UID: int32(id),
	}
	data = model.GetSth(data)
	SendResponse(r, nil, gin.H{
		"info": data,
	})
}

