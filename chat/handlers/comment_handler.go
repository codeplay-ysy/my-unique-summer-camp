package handlers

import (
	"chat/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 创建评论处理器
func CreateCommentHandler(c *gin.Context) {
	// 验证登录状态
	userID, ok := models.GetLoggedInUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "请先登录",
		})
		return
	}

	// 获取请求参数
	postIDStr := c.PostForm("post_id")
	content := c.PostForm("content")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "该帖子不存在",
		})
		return
	}
	// 生成匿名
	anonymousName, _ := models.GenerateAnonymousName(userID, uint(postID))

	// 创建评论
	err = models.CreateComment(uint(postID), userID, content, anonymousName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "创建评论失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "评论发布成功",
	})
}

// 处理创建评论的评论
func CreateCommentReplyHandler(c *gin.Context) {
	// 验证登录状态
	userID, ok := models.GetLoggedInUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "请先登录",
		})
		return
	}

	// 解析请求中的参数
	commentIDStr := c.PostForm("commentID")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "该评论不存在",
		})
		return
	}
	content := c.PostForm("content")
	postIDStr := c.PostForm("postID")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "该帖子不存在",
		})
		anonymousName, _ := models.GenerateAnonymousName(userID, uint(postID))
		// 调用创建评论的评论的函数
		err = models.CreateCommentReply(uint(commentID), userID, content, anonymousName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "回复失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "回复成功",
		})
	}
}
