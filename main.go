package main

import (
	_ "main/docs"
	"main/router"
	"main/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Melting API
//	@description	Backend system of Muxi_Melting
//	@version		1.0
//	@contact.name	@a48zhang & @Cg1028
//	@contact.email	3557695455@qq.com 2194028175@qq.com
//	@schemes		http
//	@BasePath		/api/v1
func main() {
	service.Init()
	r := gin.Default()
	router.Register(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":65000")
}
