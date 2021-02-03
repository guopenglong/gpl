package main

import (
	"ginlesson/controllar"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	DB *gorm.DB
)


func initMySQL()(err error)  {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB,err = gorm.Open("mysql",dsn)
	if err !=nil{
		return
	}
	return DB.DB().Ping()
}


func main() {
	//创建数据库

	//连接数据库
	err := initMySQL()
	if err!=nil{
		panic(err)
	}
	defer DB.Close() //程序退出
	//模型绑定

	r:=gin.Default()
	r.Static("/static","static")
	r.LoadHTMLGlob("templates/*")
	 r.GET("/",controller.IndexHandler)
	//v1
	v1Group :=r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo",controller.CreateTodo)
		//查看所有待办
		v1Group.GET("/todo", controller.GetTodoList)
		//查看某一个待办
		v1Group.GET("/todo/:id", func(c *gin.Context) {
		})
		//修改
		v1Group.PUT("todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("todo/:id",controller.DeleteATodo)
	}



	r.Run(":9090")
}
