package router

import (
	"main/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {

	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})
	users := e.Group("/users")
	{
		users.POST("/login", handler.Login)    //登录 ok
		users.GET("/", handler.GetUserInfo)    //获取用户信息 ok
		users.POST("/", handler.UploadProfile) //上传更新 ok
		
	}
	projects := users.Group("/projects")
	{
		projects.GET("/", handler.Getprojects) //获取项目信息 ok
		projects.POST("/", handler.Upproject) //更新项目 ok

	}
	proshare := e.Group("/share")
	{
		proshare.GET("/", handler.Getproject)    //获取分享项目 ok
		proshare.POST("/share", handler.Uptables) //上传“我”的信息 ok

	}
	project := users.Group("/project")
	{
		project.GET("/", handler.Getproject) //获取项目信息 ok
	}

} 
