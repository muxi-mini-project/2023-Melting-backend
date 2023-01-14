package handler

import (
	"main/model"
	"main/model/db"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadProfile godoc
//
//	@Summary		upload profile
//	@Tags			user
//	@Description	upload sth with its UserID
//	@Accept			application/json
//	@Produce		json
//	@Param			Authorization	header		string	true	"token"
//	@Param			Profile			body		db.User	true	"new user profile"
//	@Success		200				{string}	string
//	@Failure		500				{object}	handler.Response	"update failed"
//	@Router			/users [put]
func UploadProfile(r *gin.Context) {
	id := r.GetInt("userID")
	data := db.User{
		UID: int32(id),
	}
	err := r.ShouldBindJSON(&data)
	if err != nil {
		SendError(r, err, data, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	if err := model.UpdateSth(data); err != nil {
		SendError(r, err, data, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, model.NoResponse)
}

// UploadPhoto godoc
//
//	@Summary		upload photo
//	@Tags			user
//	@Description	upload photo with its UserID
//	@Accept			image/jpeg
//	@Produce		json
//	@Param			file			formData	object	true	"the photo of the user"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	string
//	@Failure		400				{object}	handler.Response	"file not received"
//	@Failure		500				{object}	handler.Response	"failed to upload file"
//	@Router			/users/photo [put]
func UploadPhoto(r *gin.Context) {
	id := r.GetInt("userID")
	H, err := r.FormFile("file")
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
	}
	file, err := H.Open() // Warning: file must be *.jpg
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
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
	err = model.UpdateSth(data)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusInternalServerError)
	}
	SendResponse(r, data)
}

// GetUserInfo godoc
//
//	@Summary		Get User's info
//	@Tags			user
//	@Description	Get User's info with its userID
//	@Param			Authorization	header	string	true	"token"
//	@Produce		json
//	@Success		200	{object}	db.User
//	@Router			/users [get]
func GetUserInfo(r *gin.Context) {
	id := r.GetInt("userID")
	data := db.User{
		UID: int32(id),
	}
	data = model.GetSth(data)
	SendResponse(r, data)
}
