package main

import (
	"chat/handlers"
	"chat/middlewares"
	"chat/models"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func init() {
	// 连接到 Redis
	models.RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 服务器地址和端口
	})

	// 检查 Redis 连接是否成功
	_, err := models.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("无法连接到 Redis：%v", err)
	}

	models.Db, err = gorm.Open("postgres", "user=postgres dbname=chat password=123000 sslmode=disable")
	if err != nil {
		log.Fatalf("无法连接到数据库：%v", err)
	}

	// 自动创建 User、Post 和 Comment 表
	models.Db.AutoMigrate(&models.User{}, &models.Post{})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// 将中间件类型转换为 gin.HandlerFunc
	ensureLoggedInMiddleware := middlewares.EnsureLoggedIn()
	// 注册中间件
	r.Use(ensureLoggedInMiddleware)

	// 用户认证和授权
	r.POST("/register", handlers.RegisterUserHandler)
	r.POST("/login", handlers.LoginUserHandler)
	r.POST("/password/reset", handlers.GeneratePasswordResetCodeHandler)
	r.POST("/password/verify", handlers.VerifyPasswordResetCodeHandler)

	// 帖子操作
	r.POST("/post", handlers.CreatePostHandler)
	r.PUT("/post/:id", handlers.UpdatePostHandler)
	r.DELETE("/post/:id", handlers.DeletePostHandler)
	r.GET("/post/:id", handlers.GetPostHandler)
	r.GET("/post/hot", handlers.GetHotPostHandler)
	r.GET("/post/search/keyword/:keyword", handlers.SearchPostsByKeywordHandler)
	r.GET("/post/search/time/:time", handlers.SearchPostsByCreateTimeHandler)

	// 评论操作
	r.POST("/post/:id/comment", handlers.CreateCommentHandler)
	r.POST("/comment/:id/reply", handlers.CreateCommentReplyHandler)

	// 举报帖子
	r.POST("/post/:id/report", handlers.ReportPostHandler)

	// 生成分享链接
	r.POST("/post/:id/share", handlers.ShareLinkHandler)

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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

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

	wg.Wait()
}
