package service

import "github.com/gin-gonic/gin"

func Auth(c *gin.Context) {
	ctx, err := auth.ParseRequest(c)
	if err != nil {

		c.Abort()
		return
	}

	c.Set("userID", ctx.ID)
	c.Set("expiresAt", ctx.ExpiresAt)

	c.Next()
}
