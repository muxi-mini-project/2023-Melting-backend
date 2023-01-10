package service

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// 声明一个token容器
var Signedkey = []byte("Muxi_Melting")

// 定义自己的声明结构体
type Myclaims struct {
	UID int
	jwt.StandardClaims
}

// 生成jwt加密的token
func Newtoken(UID int) (string, error) {
	claims := Myclaims{
		UID,
		jwt.StandardClaims{
			Issuer:    "Muxi_Melting",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Signedkey)
}

// token解析函数，配合GetToken工作
func Parsetoken(tokenString string) (*jwt.Token, *Myclaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Myclaims{},
		func(token *jwt.Token) (interface{}, error) {
			return Signedkey, nil
		})
	cl, _ := token.Claims.(*Myclaims)
	return token, cl, err
}

/* // 获取并解析token，返回带在token中的UID
func GetToken(r *gin.Context) {
	tokenString := r.GetHeader("Authorization")
	if tokenString == "" {
		r.AbortWithError(403, model.ErrAuthInvalid)
		return
	}
	token, claims, err := Parsetoken(tokenString)
	if err != nil || !token.Valid {
		r.AbortWithError(403, model.ErrAuthInvalid)
		return
	}
	r.Set("userID", claims.UID)
	r.Set("expiresAt", claims.ExpiresAt)
	r.Next()
} */
