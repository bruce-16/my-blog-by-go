package main

import (
	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zachrey/blog/routers"
)

func main() {
	r := gin.New()
	// 设置跨域
	r.Use(cors.Default())

	// 设置日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 使用日志中间件
	r.Use(gin.Logger())
	// 设置静态文件夹
	r.Static("/static", "./static")
	// 加载路由
	routers.LoadRouters(r)
	r.Run(":8888")
}
