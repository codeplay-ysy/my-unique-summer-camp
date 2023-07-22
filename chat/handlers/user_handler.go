// user_handler.go

package handlers

import (
	"chat/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	genderStr := c.PostForm("gender")

	// 检查 gender 是否有效（仅支持 Male 和 Female）
	gender := models.Gender(genderStr)
	if gender != models.Male && gender != models.Female {
		c.JSON(400, gin.H{
			"error": "无效的性别选项",
		})
		return
	}

	// 调用注册函数
	err := models.RegisterUser(name, email, password, gender)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
func LoginUserHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	// 调用登录函数并获取用户 ID
	userID, err := models.LoginUser(email, password)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 设置用户 ID 到 Cookie
	c.SetCookie("user_id", string(userID), 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"message": "登录成功",
	})
}
func ReadPostHandler(c *gin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的帖子ID",
		})
		return
	}

	post, err := models.ReadPost(uint(postID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "帖子不存在",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}
