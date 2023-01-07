package handler

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"main/model/db"
	"main/service"
)

// UploadProfile godoc
// @Summary upload profile
// @Tags manipulate
// @Description upload sth with its UserID
// @Accept application/json
// @Produce json
// @Param userID query string true "the userID of the user"
// @Param token header string true "token"
// @Success 200 {string} "upload finished"
// @Router api/v1 [post]
func UploadProfile(r *gin.Context) {
	id := r.GetInt("userID")
	data := db.User{
		UID: int32(id),
	}
	model.UpdateSth(data)
	SendResponse(r, nil, model.NoResponse)
}

// UploadPhoto godoc
// @Summary upload photo
// @Tags manipulate
// @Description upload photo with its UserID
// @Accept multipart/form-data
// @Produce json
// @Param formData form object true "the photo of the user"
// @Param token header string true "token"
// @Success 200 {object} "received data"
// @Router api/v1/photo [post]
func UploadPhoto(r *gin.Context) {
	id := r.GetInt("userID")
	H, err := r.FormFile("file")
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), 403)
	}
	file, err := H.Open() // Warning: file must be *.jpg
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), 403)
	}
	url, err := service.UploadProfilePhoto(id, file, H.Size)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), 500)
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
// @Summary Get User's info
// @Description Get User's info with its userID
// @Param id query string true "the  of the user"
// @Param token header string true "token"
// @Produce json
// @Success 200 {object} db.Userdb.ProposalInfo
// @Failure 404 {object} handler.Response
// @Router /api/v1 [get]
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

// TODO
func Uptables(r *gin.Context) {
	id := r.GetInt("InfoID")
	tables := r.GetString("tables")
	data := db.ProposalInfo{
		InfoID:         int32(id),
		OptionalTables: string(tables),
	}
	data = model.GetSth(data)
	r.JSON(200, "ok")
}
