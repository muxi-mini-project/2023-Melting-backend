package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"time"
)

var Signedkey = []byte("Muxi_Melting")

func Newtoken(UID int) (string, error) {
	claims := jwt.StandardClaims{
		Issuer:    "Muxi_Melting",
		IssuedAt:  time.Now().Unix,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Signedkey)
}

func GetToken(r *gin.Context) {
	header := r.GetHeader("Authorization")
	if header == "" {
		return
	}

}
