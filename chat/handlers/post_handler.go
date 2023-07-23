package handlers

import (
	"fmt"
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
	authorID, ok := models.GetLoggedInUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "请先登录",
		})
		return
	}

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
	authorID, ok := models.GetLoggedInUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "请先登录",
		})
		return
	}

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
	authorID, ok := models.GetLoggedInUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "请先登录",
		})
		return
	}

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

// 搜索帖子处理器（关键词获取）
func SearchPostsByKeywordHandler(c *gin.Context) {
	keyword := c.Query("keyword")
	posts, err := models.SearchPostsByKeyword(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取帖子失败",
		})
		return
	}

	// 只保留需要的字段：标题、创建时间、总浏览量
	var filteredPosts []gin.H
	for _, post := range posts {
		filteredPost := gin.H{
			"title":       post.Title,
			"created_at":  post.CreatedAt,
			"total_views": post.TotalViews,
		}
		filteredPosts = append(filteredPosts, filteredPost)
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": filteredPosts,
	})
}

// 搜索帖子处理器（创建时间获取）
func SearchPostsByCreateTimeHandler(c *gin.Context) {
	createTimeStr := c.Query("create_time")
	createTime, err := time.Parse("2006-01-02", createTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的创建时间格式，正确格式为：YYYY-MM-DD",
		})
		return
	}

	posts, err := models.SearchPostsByCreateTime(createTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取帖子失败",
		})
		return
	}

	// 只保留需要的字段：标题、创建时间、总浏览量
	var filteredPosts []gin.H
	for _, post := range posts {
		filteredPost := gin.H{
			"title":       post.Title,
			"created_at":  post.CreatedAt,
			"total_views": post.TotalViews,
		}
		filteredPosts = append(filteredPosts, filteredPost)
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": filteredPosts,
	})
}

// 获取帖子及其评论处理器
func GetPostHandler(c *gin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的帖子ID",
		})
		return
	}

	// 获取帖子信息
	post, err := models.GetPostByID(uint(postID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "帖子不存在",
		})
		return
	}

	// 获取帖子下的所有评论
	comments, err := models.GetCommentsByPostID(uint(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取评论",
		})
		return
	}

	// 遍历评论，筛选出匿名评论和匿名评论的评论
	var anonymousComments []models.Comment
	for _, comment := range comments {
		if comment.AnonymousName != "" {
			// 添加匿名评论
			anonymousComments = append(anonymousComments, comment)

			// 遍历匿名评论的评论
			for _, reply := range comment.Replies {
				if reply.AnonymousName != "" {
					// 添加匿名评论的评论
					anonymousComments = append(anonymousComments, reply)
				}
			}
		}
	}

	// 清空评论和评论的评论，只返回匿名评论和匿名评论的评论的内容和匿名名称
	post.Comments = nil
	post.Comments = anonymousComments

	c.JSON(http.StatusOK, post)
}

// 处理分享帖子的请求
func ShareLinkHandler(c *gin.Context) {
	shareLink := c.Param("shareLink")

	// 根据分享链接查找帖子
	post, err := models.GetPostByShareLink(shareLink)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "帖子不存在",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

// 处理举报帖子的请求
func ReportPostHandler(c *gin.Context) {
	// 获取举报信息
	postIDStr := c.Param("id")
	reason := c.PostForm("reason")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的帖子ID",
		})
		return
	}
	// 获取当前登录用户的ID
	userID, ok := models.GetLoggedInUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "请先登录",
		})
		return
	}

	// 将举报信息保存到数据库
	err = models.CreateReport(uint(postID), userID, reason)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("举报帖子失败：%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功举报帖子",
	})
}

// 处理获取热门帖子的请求
func GetHotPostHandler(c *gin.Context) {
	// 调用获取热门帖子的函数
	posts, err := models.GetHotPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取热门帖子",
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}
