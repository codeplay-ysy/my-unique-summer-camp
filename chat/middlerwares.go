package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("user_id")
		if err != nil || cookie == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "未登录",
			})
			c.Abort()
			return
		}

		userId := extractUserIdFromCookie(cookie)
		if userId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "未登录",
			})
			c.Abort()
			return
		}

		// 在上下文中保存用户ID，方便后续的处理函数使用
		c.Set("user_id", userId)

		// 继续处理后续的请求
		c.Next()
	}
}

// 从cookie中提取用户ID
func extractUserIdFromCookie(cookieValue string) string {
	if strings.HasPrefix(cookieValue, "user_id=") {
		return cookieValue[8:]
	}
	return ""
}
