package handler

import (
	"main/model"
	"main/service"

	"github.com/gin-gonic/gin"
)

func UploadProfile(r *gin.Context) {
	id := r.GetInt("userID")
	data := model.User{
		UID: int32(id),
	}
	model.UpdateSth(&data)
}

func UploadPhoto(r *gin.Context) {
	var data model.User
	r.ShouldBindJSON(&data)
	H, _ := r.FormFile("file")
	file, _ := H.Open() // Warning: file must be *.jpg
	service.UploadProfilePhoto(int(data.UID), file, H.Size)
}

func GetUserInfo(r *gin.Context) {
	id := r.GetInt("userID")
	data := model.User{
		UID: int32(id),
	}
	model.GetSth(&data)
	r.JSON(200, data)
}

func Updraft(r *gin.Context) {

}

func Getdraft(r *gin.Context) {

}

func Getproject(r *gin.Context) {

}

func Upproject(r *gin.Context) {

}

func Getprojects(r *gin.Context) {

}

func Uptables(r *gin.Context) {

}
