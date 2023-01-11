package middleware

import (
	"main/model"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenParser(r *gin.Context) {
	tokenString := r.GetHeader("Authorization")
	if tokenString == "" {
		r.AbortWithError(http.StatusForbidden, model.ErrAuthInvalid)
		return
	}
	token, claims, err := service.Parsetoken(tokenString)
	if err != nil || !token.Valid {
		r.AbortWithError(http.StatusForbidden, model.ErrAuthInvalid)
		return
	}
	r.Set("userID", claims.UID)
	r.Set("expiresAt", claims.ExpiresAt)
	r.Next()
}
