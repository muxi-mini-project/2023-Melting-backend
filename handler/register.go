package handler

import (
	"main/model"
	"main/model/db"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
//
//	@Summary		register account
//	@Tags			register and login
//	@Description	create a new account in certain way
//	@Param			type	query	string	false	"the type of the register"
//	@Param			data	body	db.User	true	"register data"
//	@Produce		json
//	@Success		200	{object}	db.User
//	@Failure		400	{object}	handler.Response	"param not satisfied"
//	@Router			/register [post]
func Register(r *gin.Context) {
	typ := r.Query("type")
	var data db.User
	err := r.ShouldBindJSON(&data)
	if err != nil {
		SendError(r, err, data, model.ErrorSender(), http.StatusBadRequest)
		return
	}
	switch typ {
	case "MuxiPass":
		//TODO
	case "ccnu":
		//TODO
	case "qq":
		err = service.CreateWithQQ(data)
	default:
		err = service.CreateNative(data)
	}
	if err != nil {
		SendError(r, err, data, model.ErrorSender(), http.StatusForbidden)
		return
	}
	SendResponse(r, data)
}
