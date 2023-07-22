package handlers

import (
	"net/http"
	"strconv"

	"chat/models"
	"time"

	"github.com/gin-gonic/gin"
)

// 发帖处理器
func CreatePostHandler(c *gin.Context) {
	// 获取用户提交的数据
	title := c.PostForm("title")
	content := c.PostForm("content")

	// 获取当前登录用户的 ID
	authorID := GetLoggedInUserID(c)

	// 调用模型函数创建帖子
	err := models.CreatePost(title, content, authorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "发帖失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "发帖成功",
	})
}

// 删帖处理器
func DeletePostHandler(c *gin.Context) {
	// 获取帖子ID参数
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的帖子ID",
		})
		return
	}

	// 获取当前登录用户的 ID
	authorID := GetLoggedInUserID(c)

	// 调用模型函数删除帖子
	err = models.DeletePost(uint(postID), authorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// 更新帖子处理器
func UpdatePostHandler(c *gin.Context) {
	// 获取帖子ID参数
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的帖子ID",
		})
		return
	}

	// 获取用户提交的数据
	title := c.PostForm("title")
	content := c.PostForm("content")

	// 获取当前登录用户的 ID
	authorID := GetLoggedInUserID(c)

	// 调用模型函数更新帖子
	err = models.UpdatePost(uint(postID), authorID, title, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

func GetLoggedInUserID(c *gin.Context) uint {

	userID, _ := c.Get("userID")
	return userID.(uint)
}

// 读取帖子处理器（关键词获取）
func SearchPostsByKeywordHandler(c *gin.Context) {
	keyword := c.Query("keyword")
	posts, err := models.SearchPostsByKeyword(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取帖子失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// 读取帖子处理器（创建时间获取）
func GetPostsByCreateTimeHandler(c *gin.Context) {
	createTimeStr := c.Query("create_time")
	createTime, err := time.Parse("2006-01-02", createTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的创建时间格式，正确格式为：YYYY-MM-DD",
		})
		return
	}

	posts, err := models.GetPostsByCreateTime(createTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取帖子失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
