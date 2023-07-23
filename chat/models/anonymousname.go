package models

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

type AnonymousName struct {
	UserID    uint
	PostID    uint
	Anonymous string
}

// 查询数据库中是否存在对应的匿名名称
func GenerateAnonymousName(userID, postID uint) (string, error) {
	var anonymousName AnonymousName
	err := Db.Where("user_id = ? AND post_id = ?", userID, postID).First(&anonymousName).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// 不存在匿名名称，生成一个新的匿名名称
			anonymous := generateUniqueAnonymousName()

			// 将新的匿名名称保存到数据库
			anonymousName = AnonymousName{
				UserID:    userID,
				PostID:    postID,
				Anonymous: anonymous,
			}
			err := Db.Create(&anonymousName).Error
			if err != nil {
				return "", err
			}

			// 返回新生成的匿名名称
			return anonymous, nil
		}

		// 查询过程中发生了其他错误
		return "", err
	}

	// 存在匿名名称，直接返回
	return anonymousName.Anonymous, nil
}

var fruits = []string{
	"苹果", "香蕉", "橙子", "草莓", "葡萄",
	"西瓜", "樱桃", "梨子", "桃子", "柠檬",
	"橘子", "芒果", "蓝莓", "菠萝", "木瓜",
	"柚子", "哈密瓜", "猕猴桃", "杨梅", "荔枝",
}

func generateUniqueAnonymousName() string {
	// 设置随机种子，防止每次生成的随机数相同
	rand.Seed(time.Now().UnixNano())

	// 随机选择一个水果
	fruit := fruits[rand.Intn(len(fruits))]

	anonymousName := fruit

	return anonymousName
}
