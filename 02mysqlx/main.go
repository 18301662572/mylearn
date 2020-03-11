package main

import (
	"code.oldbody.com/studygolang/mylearn/02mysqlx/db"
	"code.oldbody.com/studygolang/mylearn/02mysqlx/handlers"
	"github.com/gin-gonic/gin"
)

//user管理页面

func main() {
	//程序启动就要应该连接数据库
	err := db.InitDB()
	if err != nil {
		panic(err)
	}
	//使用gin框架
	r := gin.Default()
	//加载页面
	r.LoadHTMLGlob("templates/*")
	//加载静态页面 (代码里使用的路径，实际保存静态文件的路径)
	r.Static("/user/jtym", "statics")

	////查看用户列表(路由，handler)，使用get方法
	//r.GET("/user/list", handlers.UserListHandler)
	////添加用户(路由，handler)，使用Any(),通过c.Request.Method判断
	//r.Any("/user/add", handlers.AddUserAnyHandler)
	////r.GET("/user/add", handlers.AddUserHandler)
	////r.POST("/user/add", handlers.CreateUserHandler)
	////修改用户
	//r.GET("/user/update", handlers.UpdateUserHandler)
	//r.POST("/user/update", handlers.UpdateUserPostHandler)
	////删除用户信息
	//r.GET("/user/delete", handlers.DeleteUserHandler)

	//路由组
	userGroup := r.Group("/user")
	{
		userGroup.GET("/list", handlers.UserListHandler)
		userGroup.Any("/add", handlers.AddUserAnyHandler)
		//userGroup.GET("/add", handlers.AddUserHandler)
		//userGroup.POST("/add", handlers.CreateUserHandler)
		userGroup.GET("/update", handlers.UpdateUserHandler)
		userGroup.POST("/update", handlers.UpdateUserPostHandler)
		userGroup.GET("/delete", handlers.DeleteUserHandler)
	}

	//启动路由
	r.Run(":9091")
}
