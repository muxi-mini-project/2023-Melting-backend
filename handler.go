package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 新建用户
func (User) creat_newuser(operating_user) {
	if operating_user != nil {
		db.Create(&operating_user)
	}
}

// 新建项目
func (project) creat_newproject(operating_project) {
	if operating_project != nil {
		db.Create(&operating_project)
	}
}

// 新建活动
func (activity) creat_newactivity(operating_activity) {
	if operating_activity != nil {
		db.Create(&operating_activity)
	}
}

// 添加游戏题库
func (question) creat_newquestion(operating_question) {
	if operating_question != nil {
		db.Create(&operating_question)
	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "搭建完成")
	})
	r.Run(":8888") // 端口号8888
}
