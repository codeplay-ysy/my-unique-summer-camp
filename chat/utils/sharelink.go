package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateUniqueLink 生成帖子的唯一链接
func GenerateUniqueLink() string {
	// 定义链接的基本URL
	baseURL := "https://example.com/posts/"

	// 生成一个随机的8字节标识符
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// 如果出现错误，返回一个默认链接
		return baseURL + "default-link"
	}

	// 使用base64 URL编码将随机字节转换为字符串标识符
	linkIdentifier := base64.URLEncoding.EncodeToString(randomBytes)

	// 返回最终的唯一链接
	return baseURL + linkIdentifier
}
