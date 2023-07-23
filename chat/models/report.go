package models

import (
	"time"
)

type Report struct {
	ID         uint      `gorm:"primary_key"`
	PostID     uint      `gorm:"not null"` // 被举报的帖子ID
	Reason     string    `gorm:"not null"` // 举报原因
	ReporterID uint      `gorm:"not null"` // 举报人ID
	CreatedAt  time.Time `gorm:"not null"` // 举报时间
	UpdatedAt  time.Time `gorm:"not null"` // 更新时间
}

// CreateReport 创建举报
func CreateReport(postID, reporterID uint, reason string) error {
	report := Report{
		PostID:     postID,
		Reason:     reason,
		ReporterID: reporterID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := Db.Create(&report).Error
	if err != nil {
		return err
	}

	return nil
}
