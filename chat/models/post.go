package models

import (
	"chat/utils"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Post struct {
	ID            uint      `gorm:"primary_key"`
	Title         string    `gorm:"not null"`
	Content       string    `gorm:"not null"`
	AuthorID      uint      `gorm:"not null"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
	DailyViews    uint      `gorm:"not null;default:0"`
	TotalViews    uint      `gorm:"not null;default:0"`
	ShareLink     string    `gorm:"not null;unique"`
	Comments      []Comment
	AnonymousName map[string]int `gorm:"-"`
}

// 创建帖子
func CreatePost(title, content string, authorID uint) error {
	// 生成唯一的分享链接
	shareLink := utils.GenerateUniqueLink()

	post := Post{
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ShareLink: shareLink,
	}

	// 将帖子信息保存到数据库
	err := Db.Create(&post).Error
	if err != nil {
		return err
	}

	return nil
}

// 删除帖子
func DeletePost(postID, authorID uint) error {
	// 查询帖子是否存在
	var post Post
	err := Db.First(&post, postID).Error
	if err != nil {
		return err
	}

	// 检查当前登录用户是否是帖子的作者，只有作者才能删除帖子
	if post.AuthorID != authorID {
		return fmt.Errorf("没有权限删除该帖子")
	}

	// 删除帖子
	err = Db.Delete(&post).Error
	if err != nil {
		return err
	}

	return nil
}

// 更新帖子
func UpdatePost(postID, authorID uint, title, content string) error {
	// 查询帖子是否存在
	var post Post
	err := Db.First(&post, postID).Error
	if err != nil {
		return err
	}

	// 检查当前登录用户是否是帖子的作者，只有作者才能更新帖子
	if post.AuthorID != authorID {
		return fmt.Errorf("没有权限更新该帖子")
	}

	// 更新帖子内容
	post.Title = title
	post.Content = content
	post.UpdatedAt = time.Now()

	// 将更新后的帖子信息保存到数据库
	err = Db.Save(&post).Error
	if err != nil {
		return err
	}

	return nil
}

// 根据关键词搜索帖子
func SearchPostsByKeyword(keyword string) ([]Post, error) {
	var posts []Post
	err := Db.Table("posts").Select("id, title, authorid, totalviews").Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Scan(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// 根据时间搜索帖子
func SearchPostsByCreateTime(createTime time.Time) ([]Post, error) {
	var posts []Post
	err := Db.Model(&Post{}).
		Select("title, createdat, authorid, totalviews, id").
		Where("createdat = ?", createTime).
		Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// 获取前十个热门帖子
func GetHotPosts() ([]Post, error) {
	var posts []Post
	err := Db.Order("total_views desc").Limit(10).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// 根据分享链接获取帖子
func GetPostByShareLink(shareLink string) (*Post, error) {
	// 分割分享链接，获取唯一标识符部分
	parts := strings.Split(shareLink, "/")
	if len(parts) < 2 {
		return nil, errors.New("无效的分享链接")
	}

	// 获取唯一标识符
	linkIdentifier := parts[len(parts)-1]

	// 根据唯一标识符查找帖子
	var post Post
	err := Db.Where("id = ?", linkIdentifier).First(&post).Error
	if err != nil {
		return nil, err
	}

	// 增加帖子的总浏览量
	post.TotalViews++
	Db.Save(&post)

	return &post, nil
}

// 根据ID读取指定帖子
func GetPostByID(postID uint) (*Post, error) {
	var post Post
	err := Db.First(&post, postID).Error
	if err != nil {
		return nil, err
	}
	//浏览量增加
	post.DailyViews++
	post.TotalViews++
	Db.Save(&post)
	return &post, nil
}

/*
更新每日浏览量（考虑了两种方案。一种是专门设置PostViews结构体，这样实现了数据分离
不容易导致数据冗余，也不会因为并发写入而导致数据不一致。另一种是直接在Post结构体中
添加DailyViews和TotalViews字段，这样简化了数据结构，方便查询，还减少了表的数量。
但是这样可能会导致数据冗余，而且并发写入时可能会导致数据不一致。）
*/
func UpdateDailyViews() {
	// 获取当前日期的零点时间
	now := time.Now()
	year, month, day := now.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	// 查询所有帖子
	var posts []Post
	Db.Find(&posts)

	// 遍历所有帖子，更新每日浏览量并重置每日浏览量为0
	for _, post := range posts {
		// 如果上次更新时间不是今天，则重置每日浏览量为0
		if post.UpdatedAt.Before(startOfDay) {
			post.DailyViews = 0
		}

		// 更新总浏览量，并保存到数据库
		post.TotalViews += post.DailyViews
		Db.Save(&post)
	}
}
