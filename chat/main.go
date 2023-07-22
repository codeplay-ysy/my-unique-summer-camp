package main

import (
	"chat/handlers"
	"chat/models"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func init() {
	var err error
	models.Db, err = gorm.Open("postgres", "user=postgres dbname=chat password=123000 sslmode=disable")
	if err != nil {
		log.Fatalf("无法连接到数据库：%v", err)
	}

	// 自动创建 User、Post 和 Comment 表
	models.Db.AutoMigrate(&models.User{}, &models.Post{})
}
func setupRouter() *gin.Engine {
	r := gin.Default()

	// 路由组1：用户相关
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handlers.RegisterUserHandler)
		userGroup.POST("/login", handlers.LoginUserHandler)
		// 添加其他用户相关的路由处理函数
	}

	// 路由组2：帖子相关
	postGroup := r.Group("/post")
	{
		postGroup.POST("/create", handlers.CreatePostHandler)
		postGroup.DELETE("/delete/:id", handlers.DeletePostHandler)
		postGroup.PUT("/update/:id", handlers.UpdatePostHandler)
		postGroup.GET("/read/:id", handlers.ReadPostHandler)
		// 添加其他帖子相关的路由处理函数
	}

	// 添加其他路由组和路由处理函数...

	return r
}
func main() {
	defer models.Db.Close()

	// 初始化数据库连接和路由
	r := setupRouter()

	// 启动 HTTP 服务器
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("HTTP 服务器启动失败：%v", err)
	}
	// 启动定时器，每天凌晨0点更新每日浏览量
	go func() {
		for {
			// 获取当前时间
			now := time.Now()

			// 计算下一个零点时间
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, now.Location())

			// 等待到下一个零点时间执行更新
			t := time.NewTimer(next.Sub(now))
			<-t.C

			// 更新每日浏览量
			models.UpdateDailyViews()
		}
	}()
}
