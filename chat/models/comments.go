package models

import "time"

type Comment struct {
	ID            uint      `gorm:"primary_key"`
	Content       string    `gorm:"not null"`
	PostID        uint      `gorm:"not null"` // 所属帖子ID
	AuthorID      uint      `gorm:"not null"` // 评论作者ID
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
	ParentID      uint      // 父评论ID
	Replies       []Comment // 子评论
	AnonymousName string    // 匿名名称
}

// 创建评论
func CreateComment(postID, authorID uint, content string, anonymousName string) error {
	comment := Comment{
		Content:       content,
		PostID:        postID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		AnonymousName: anonymousName,
	}

	// 将评论保存到数据库
	err := Db.Create(&comment).Error
	if err != nil {
		return err
	}

	return nil
}

// 创建评论的评论
func CreateCommentReply(commentID, authorID uint, content, anonymousName string) error {
	comment := Comment{
		Content:       content,
		PostID:        0, // 这里设为0，表示这是评论的评论，没有对应的帖子ID
		ParentID:      commentID,
		AuthorID:      authorID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		AnonymousName: anonymousName,
	}

	// 将评论的评论保存到数据库
	err := Db.Create(&comment).Error
	if err != nil {
		return err
	}

	return nil
}

// 获取帖子下的所有评论
func GetCommentsByPostID(postID uint) ([]Comment, error) {
	var comments []Comment
	err := Db.Where("post_id = ? AND parent_id IS NULL", postID).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	// 递归获取每个评论下的回复
	for i := range comments {
		err := getReplies(&comments[i])
		if err != nil {
			return nil, err
		}
	}

	return comments, nil
}

func getReplies(comment *Comment) error {
	var replies []Comment
	err := Db.Where("parent_id = ?", comment.ID).Find(&replies).Error
	if err != nil {
		return err
	}

	// 递归获取每个回复的回复
	for i := range replies {
		err := getReplies(&replies[i])
		if err != nil {
			return err
		}
	}

	comment.Replies = replies
	return nil
}
