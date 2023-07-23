package models

import (
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var RedisClient *redis.Client

var ctx = context.Background()

const ResetCodeInterval = 60 // 60秒，即允许每隔60秒重新发送验证码

// CreatePasswordResetCode 创建找回密码时的验证码
func CreatePasswordResetCode(email, code string, createTime, expireTime time.Time) error {
	// 将验证码信息存储到Redis
	err := RedisClient.Set(ctx, "reset_code:"+email, code, expireTime.Sub(createTime)).Err()
	if err != nil {
		return err
	}

	return nil
}

// 校验找回密码时生成的验证码
func VerifyPasswordResetCode(email, code string) error {
	// 从Redis中获取验证码
	storedCode, err := RedisClient.Get(ctx, "reset_code:"+email).Result()
	if err != nil {
		if err == redis.Nil {
			return errors.New("验证码错误")
		}
		return err
	}

	// 校验验证码
	if storedCode != code {
		return errors.New("验证码错误")
	}

	// 从Redis中删除验证码
	RedisClient.Del(ctx, "reset_code:"+email)

	return nil
}

// 检查是否允许重新发送验证码
func IsCodeGenerationAllowed(email string) bool {
	// 从Redis中获取上一次获取验证码的时间戳
	lastCodeTime, err := RedisClient.Get(ctx, "last_code_time:"+email).Int64()
	if err != nil && err != redis.Nil {
		// 如果发生错误，并且不是Key不存在的错误，直接返回true，允许生成验证码
		return true
	}

	// 获取当前时间戳
	currentTime := time.Now().Unix()

	// 判断是否允许重新发送验证码
	if currentTime-lastCodeTime < ResetCodeInterval {
		// 在指定的时间间隔内已经获取过验证码，不允许重新发送
		return false
	}

	// 允许重新发送验证码
	return true
}
