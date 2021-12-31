package lib

import (
	appConfig "nightingale/config"
	"net/smtp"
)

// 设置邮件服务器账号
var (
	SMTPHost     = appConfig.EmailConfig["smtpHost"]
	SMTPPort     = ":" + appConfig.EmailConfig["smtpPort"]
	SMTPUsername = appConfig.EmailConfig["smtpUser"]
	SMTPPassword = appConfig.EmailConfig["smtpPass"]
)
// 纯文本内容邮件函数
func SendEmail(receiver,subject,body string) error {
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	msg := []byte("Subject: "+ subject +"\r\n\r\n" + body + "\r\n")
	return smtp.SendMail(SMTPHost+SMTPPort, auth, SMTPUsername, []string{receiver}, msg)
}
// HTML内容邮件函数
func SendHTMLEmail(receiver,subject string, html []byte) error {
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	msg := append([]byte("Subject:"+ subject +"\r\n"+
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"),
		html...)
	return smtp.SendMail(SMTPHost+SMTPPort, auth, SMTPUsername, []string{receiver}, msg)
}
