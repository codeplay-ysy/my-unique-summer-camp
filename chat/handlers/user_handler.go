// user_handler.go

package handlers

import (
	"chat/models"
	"chat/utils"
	"net/http"
	"time"

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
	c.SetCookie("user_id", "user_id="+string(userID), 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"message": "登录成功",
	})
}

// 生成找回密码时的验证码处理器
func GeneratePasswordResetCodeHandler(c *gin.Context) {
	// 获取用户提交的数据
	email := c.PostForm("email")

	// 检查是否已经获取过验证码，如果在60秒内重新获取，则返回错误提示
	if models.IsCodeGenerationAllowed(email) {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": "验证码获取过于频繁，请稍后再试",
		})
		return
	}

	// 生成随机的验证码
	code, err := utils.GenerateRandomCode(4)

	// 设置验证码有效期为10分钟
	expireTime := time.Now().Add(10 * time.Minute)

	// 保存验证码到Redis
	err = models.CreatePasswordResetCode(email, code, time.Now(), expireTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法生成验证码",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "验证码已发送，请注意查收",
	})
}

// 校验找回密码时的验证码处理器
func VerifyPasswordResetCodeHandler(c *gin.Context) {
	// 获取用户提交的数据
	email := c.PostForm("email")
	code := c.PostForm("code")

	// 校验验证码
	err := models.VerifyPasswordResetCode(email, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "验证码验证通过",
	})
}
