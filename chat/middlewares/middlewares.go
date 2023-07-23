package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的路径
		path := c.Request.URL.Path

		// 判断是否是需要登录权限的处理器，这里可以根据实际情况自定义需要登录权限的处理器
		requiresLogin := path == "/post" || path == "/post/:id" || path == "/post/:id/comment" || path == "/comment/:id/reply" || path == "/post/:id/report"

		// 如果需要登录权限且用户未登录，则中止请求并返回错误响应
		if requiresLogin {
			cookie, err := c.Cookie("user_id")
			if err != nil || cookie == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "请先登录",
				})
				c.Abort()
				return
			}
		}

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
