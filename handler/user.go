package handler

import (
	"main/model"
	"main/service"

	"github.com/gin-gonic/gin"
)

func UploadProfile(r *gin.Context) {
	var data model.User
	// TODO
	r.ShouldBindJSON()
}

func UploadPhoto(r *gin.Context) {
	var data model.User
	r.ShouldBindJSON(&data)
	H, _ := r.FormFile("file")
	file, _ := H.Open() // Warning: file must be *.jpg
	service.UploadProfilePhoto(int(data.UID), file, H.Size)
}

func GetUserInfo(r *gin.Context) {
	token, err := r.Cookie("token")
	if err != nil {
		//TODO go to the auth middleware
	}
}
