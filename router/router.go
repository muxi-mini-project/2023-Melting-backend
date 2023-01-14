package router

import (
	/* "html/template" 完美*/
	"main/handler"
	"main/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})

	v1 := e.Group("/api/v1")
	{
		v1.POST("/login", handler.Login)       // test pass
		v1.POST("/register", handler.Register) // test pass
		v1.Use(middleware.TokenParser)
		users := v1.Group("/users")
		{
			users.GET("", handler.GetUserInfo)   //获取用户信息
			users.PUT("", handler.UploadProfile) //上传更新
			users.PUT("/photo", handler.UploadPhoto)
			users.GET("/myproject", handler.Getprojects)
		}
		project := v1.Group("/project")
		{
			project.GET("", handler.GetProject)                //获取项目信息
			project.PUT("", handler.UpdateProject)             //更新项目
			project.GET("/template", handler.GetTemplate)      //获取模板
			project.POST("/newproject", handler.CreateProject) //新建项目
			games := project.Group("/games")
			{
				games.GET("", handler.GameSelect)
				games.POST("/find", handler.FindGames)
			}
		}
	}
}
