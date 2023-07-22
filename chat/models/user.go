package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var Db *gorm.DB

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type User struct {
	Id        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	Gender    Gender `gorm:"not null"`
	CreatedAt time.Time
}

// 注册用户
func RegisterUser(name, email, password string, gender Gender) error {
	// 创建用户结构
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := User{
		Name:      name,
		Email:     email,
		Password:  string(hashedPassword),
		Gender:    gender,
		CreatedAt: time.Now(),
	}

	// 将用户信息保存到数据库
	err = Db.Create(&user).Error
	if err != nil {
		return fmt.Errorf("注册失败：%v", err)
	}

	return nil
}

// 登录用户
func LoginUser(email, password string) (uint, error) {
	// 查询用户信息
	user, err := GetUserByEmail(email)
	if err != nil {
		return 0, fmt.Errorf("未查询到用户信息：%v", err)
	}
	//这里使用了哈希函数来进行密码的比对，因此实现了密码的加密
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 0, fmt.Errorf("密码错误：%v", err)
	}
	err = Db.Where(&User{Email: email, Password: password}).First(&user).Error
	if err != nil {
		return 0, fmt.Errorf("登录失败，请重试：%v", err)
	}

	return user.Id, nil
}

// 根据用户邮箱从数据库中获取用户记录（用于匹配用户密码，完成用户登录）
func GetUserByEmail(email string) (*User, error) {
	var user User
	err := Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// 未找到用户记录
			return nil, nil
		}
		// 查询过程中发生了其他错误
		return nil, err
	}
	return &user, nil
}
