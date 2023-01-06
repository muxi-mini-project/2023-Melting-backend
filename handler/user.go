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

func Getproject(r *gin.Context) {
    id:=r.GetInt("InfoID")
	data:=model.ProposalInfo{
		InfoID: int32(id),
	}
	model.GetSth(&data)
	r.JSON(200,data)
}

func Getprojects(r *gin.Context){
    id:=r.GetInt("userID")
	data,_:=model.GetProposals(id)
	r.JSON(200,data)
}

func Upproject(r *gin.Context) {
    id:=r.GetInt("InfoID")
	data:=model.ProposalInfo{
		InfoID: int32(id),
	}
	model.UpdateSth(&data)
	r.JSON(200,"ok")
}

func Uptables(r *gin.Context) {
    id:=r.GetInt("InfoID")
	tables:=r.GetString("tables")
	data:=model.ProposalInfo{
        InfoID: int32(id),
		OptionalTables:string(tables),
	}
	model.UpdateSth(&data)
	r.JSON(200,"ok")
}
