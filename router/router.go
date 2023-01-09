package router

import (
	/* "html/template" 完美*/
	"main/handler"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})

	e.POST("/login", handler.Login)       //登录 test pass
	e.POST("/register", handler.Register) // test pass

	v1 := e.Group("/api/v1")
	{
		v1.Use(service.GetToken)
		users := v1.Group("/users")
		{
			users.GET("/", handler.GetUserInfo)    //获取用户信息
			users.POST("/", handler.UploadProfile) //上传更新
			users.POST("/photo", handler.UploadPhoto)
			users.GET("/myproject", handler.Getprojects)
		}
		project := v1.Group("/project")
		{
			// TODO newproject
			project.GET("/", handler.Getproject)          //获取项目信息
			project.POST("/", handler.UpdateProject)      //更新项目
			project.GET("/template", handler.GetTemplate) //获取模板
            project.POST("/newproject",handler.CreateProject) //新建项目
			}  
	}

	/* projects := users.Group("/projects")
	{
		projects.GET("/", handler.Getprojects) //获取项目信息
	}
	proshare := e.Group("/share")
	{
		proshare.GET("/", handler.Getproject)    //获取分享项目
		proshare.POST("/", handler.Uptables) //上传“我”的信息
	}*/

}
