package main

import (
	"fmt"
	"io"
	"log"
	_ "main/docs"
	"main/router"
	"main/service"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Melting API
// @description	Backend system of Muxi_Melting
// @version		1.0
// @contact.name	@a48zhang & @Cg1028
// @contact.email	3557695455@qq.com 2194028175@qq.com
// @schemes		http
// @BasePath		/api/v1
func main() {
	logger()
	service.Init()
	r := gin.Default()
	router.Register(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":65000")
	if err != nil {
		log.Fatal(err)
	}
}

func logger() {
	y, m, d := time.Now().Date()
	target := fmt.Sprintf("./log/%v_%v_%v_%v.log", y, m, d, time.Now().Nanosecond())
	f, _ := os.Create(target)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
