package handler

import "github.com/gin-gonic/gin"

func Login(r *gin.Context) {
	typ := r.Query("type")

	switch typ {
	case "MuxiPass":
		//TODO
	case "ccnu":
		//TODO
	default:

	}
}
