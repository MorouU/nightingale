package appConfig

var EmailConfig = map[string]string{
	"smtpHost":"smtp.qq.com",
	"smtpPort": "587",
	"smtpEmailFrom":"",
	"smtpUser":"",
	"smtpPass":"",
}


var RegisterAuthData = map[string]string{

	"subject":"NI XSS攻击平台 - [账号激活]",
	"body":`
您好， %s ,请点击以下链接进行激活：%s
`,
	"urlApi" : Domain + "userEnter/api/userActive/",

}

var ForgetAuthData = map[string]string{

	"subject":"NI XSS攻击平台 - [密码找回]",
	"body":`
您好， %s ,这是您的验证码：%s ，该验证码有效期限为5分钟。
`,

}

var DefaultModuleData = map[string]string{

	"subject":"NI XSS攻击平台 - [攻击结果 - %d]",
	"body":`
您好， %s ,这是您的攻击结果：%s
`,

}