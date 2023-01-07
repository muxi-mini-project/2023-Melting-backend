package handler

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"main/model/db"
	"main/service"
	"net/http"
)

// Register godoc
// @Summary register account
// @Description create a new account with someway
// @Param type query string true "the type of the register"
// @Produce json
// @Success 200 {object} db.User
// @Router /api/register [post]
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
	SendResponse(r, nil, data)
}
