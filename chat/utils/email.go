package utils

import (
	"gopkg.in/gomail.v2"
)

// SendEmail 发送邮件，包含验证码信息
func SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "1325274509@qq.com") // 发件人邮箱，根据实际情况修改
	m.SetHeader("To", to)                    // 收件人邮箱
	m.SetHeader("Subject", subject)          // 邮件主题
	m.SetBody("text/html", body)             // 邮件正文，可以是HTML格式

	d := gomail.NewDialer("smtp.qq.com", 587, "your_email@example.com", "your_email_password") // SMTP服务器地址和端口，根据实际情况修改
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
