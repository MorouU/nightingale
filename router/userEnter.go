package appRouter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	appConfig "nightingale/config"
	lib "nightingale/lib/func"
	sqli "nightingale/lib/sql"
	"net/http"
	"strconv"
	"strings"
)

// 用户登录 GET
func pages_userEnter_login(c *gin.Context){

	// SHOW
	if c.Request.Method == "GET" {

		if (!sessionKeyExist(sessionGet(c, "loginAuthCode"), "string")) || (sessionKeyExist(sessionGet(c, "loginAuthCodeTimeOut"), "int64") && lib.CmpAuthCodeTimeOut(sessionGet(c, "loginAuthCodeTimeOut").(int64), appConfig.LoginAuthCodeConfig["timeOut"].(int))) {
			authCode, _, timeOut := func_userEnter_getLoginAuthCodeArray()
			sessionSet(c, map[interface{}]interface{}{
				"loginAuthCode":        authCode,
				"loginAuthCodeTimeOut": timeOut,
			})
		}

		authCodeB64 := func_userEnter_getLoginAuthCodeB64(sessionGet(c, "loginAuthCode").(string))

		c.HTML(
			http.StatusOK,
			"userEnter_login.ni",
			gin.H{
				"authCode": authCodeB64,
			},
		)
	}

	// LOGIN
	if c.Request.Method == "POST"{

		func_userEnter_userLogin(c)

	}
}

// 用户登录 GET
func pages_userEnter_reg(c *gin.Context){

	// SHOW
	if c.Request.Method == "GET" {

		c.HTML(
			http.StatusOK,
			"userEnter_reg.ni",
			gin.H{},
		)
	}

	// REGISTER
	if c.Request.Method == "POST" {

		func_userEnter_userRegister(c)

	}
}

// 用户密码找回 GET
func pages_userEnter_getpass(c *gin.Context){

	// SHOW
	if c.Request.Method == "GET" {
		c.HTML(
			http.StatusOK,
			"userEnter_getpass.ni",
			gin.H{},
		)
	}

	// FIND
	if c.Request.Method == "POST"{

		func_userEnter_findPass(c)

	}
}


/*
	userLogin
 */
// 验证码刷新
func api_userEnter_loginAuthCodeRefresh(c *gin.Context){

	refresh := c.DefaultPostForm("refresh", "0")

	if refresh != "1"{
		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"params error ~"},
		)
	}

	authCode,authCodeB64,timeOut := func_userEnter_getLoginAuthCodeArray()

	sessionSet(c, map[interface{}]interface{}{
		"loginAuthCode":authCode,
		"loginAuthCodeTimeOut": timeOut,
	})

	c.JSON(
		http.StatusOK,
		map[string]string{"status":"ok","contents":authCodeB64},
		)
}

// 获取验证码
func func_userEnter_getLoginAuthCodeArray()(string,string,int64){

	length := appConfig.LoginAuthCodeConfig["length"].(int)
	chars := appConfig.LoginAuthCodeConfig["chars"].(string)
	width, height := appConfig.LoginAuthCodeConfig["width"].(int), appConfig.LoginAuthCodeConfig["height"].(int)
	timeOut := lib.GetTimeUnix(false) + int64(appConfig.LoginAuthCodeConfig["timeOut"].(int))
	authCode := lib.GetRandStr(length,chars)

	return authCode , lib.BaseEn(lib.GetAuthCode(width,height,authCode),64),timeOut
}

// 获取验证码图片base64
func func_userEnter_getLoginAuthCodeB64(s string)string{
	return lib.BaseEn(lib.GetAuthCode(appConfig.LoginAuthCodeConfig["width"].(int), appConfig.LoginAuthCodeConfig["height"].(int),s),64)
}

// 检查用户名/邮箱是否唯一
func func_userEnter_checkUnique(userName,userEmail string) bool {

	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	r,_ := thisSql.Query(
			"userLogin",
			[]string{"*"},
			[]string{"userName","userEmail"},
			[]string{userName,userEmail},
			" OR ",
			"",
		)
	return len(r) > 0

}

// 用户登录 POST
func func_userEnter_userLogin(c *gin.Context){

	userAccount := c.DefaultPostForm("userAccount","")
	pass := c.DefaultPostForm("loginPass","")
	authCode := c.DefaultPostForm("authCode","")

	/*
		CHECK
	 */
	thisAuthCode := sessionGet(c,"loginAuthCode").(string)
	thisAuthCodeTime := sessionGet(c,"loginAuthCodeTimeOut").(int64)

	authCodeA,_,timeOut := func_userEnter_getLoginAuthCodeArray()

	sessionSet(c, map[interface{}]interface{}{
		"loginAuthCode":authCodeA,
		"loginAuthCodeTimeOut": timeOut,
	})

	if thisAuthCodeTime < lib.GetTimeUnix(false){
		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"Verification code timeout"},
		)

		return
	}

	if strings.ToLower(thisAuthCode) != strings.ToLower(authCode){

		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"Verification code error"},
		)
		return
	}

	/*
		SQL
	 */
	// Connect SQL
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	// Check user
	r,_ := thisSql.Query(
			"userLogin",
			[]string{"*"},
			[]string{"userName","userEmail"},
			[]string{userAccount,userAccount},
			" OR ",
			"",
		)

	if len(r) == 1{
		passSalt := r[0]["userSalt"].(string)
		passHash := lib.GetHashSignature([]byte(pass),[]byte(passSalt),"sha256",true).(string)

		if r[0]["active"].(int64) == 0{
			c.JSON(
				http.StatusOK,
				map[string]string{"status":"error","contents":"Account not enabled"},
			)
			return
		}

		if passHash != r[0]["userPasswd"].(string){
			c.JSON(
				http.StatusOK,
				map[string]string{"status":"error","contents":"Incorrect username / email or password"},
			)
			return
		}

		// OK
		authID := uuid.New()
		sessionSet(c, map[interface{}]interface{}{
			"isLogin":"1",
			"userID":r[0]["userID"].(int64),
			"userName":r[0]["userName"].(string),
			"userEmail":r[0]["userEmail"].(string),
			"authID":authID,
			"reqIP":c.ClientIP(),
		})

		// SQL Update
		rt,_ := thisSql.Update(
				"userLogin",
				[]string{"userLoginIP","lastLoginTime"},
				[]string{c.ClientIP(), lib.GetTimeDateTime()},
				[]string{"userName","userEmail"},
				[]string{r[0]["userName"].(string),r[0]["userEmail"].(string)},
				" AND ",
				"",
			)

		if rt <= 0{
			sessionClear(c)
			c.JSON(
				http.StatusOK,
				map[string]string{"status":"error","contents":"Login error"},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{"status":"ok","contents":"Login success~","uuid":authID},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		map[string]string{"status":"error","contents":"Incorrect username / email or password"},
	)

}

/*
	userRegister
 */
// 用户注册 POST
func func_userEnter_userRegister(c *gin.Context){

	userName := c.DefaultPostForm("userName","")
	email := c.DefaultPostForm("email","")
	pass := c.DefaultPostForm("pass","")
	repass := c.DefaultPostForm("repass","")
	agree := c.DefaultPostForm("agree","")
	/*
		CHECK
	 */
	// check userName
	if ! lib.CheckUsername(userName){
		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"The user name is incorrectly formatted"},
		)
		return
	}
	// check agree
	if agree != "1"{
		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"You have to agree"},
		)
		return
	}
	// cmp pass and repass
	if pass != repass{
		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"Different passwords"},
		)
		return
	}
	// verify email
	if ! lib.VerifyEmailFormat(email){
		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"Incorrect email format"},
		)
		return
	}

	/*
		SQL
	 */
	// Connect SQL
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	// Check unique
	if func_userEnter_checkUnique(userName,email){
		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"Username/email already exists"},
		)
		return
	}

	// Get user data and insert
	passSalt := uuid.New()
	passHash := lib.GetHashSignature([]byte(pass),[]byte(passSalt),"sha256",true).(string)
	activeApi := lib.GetHashSignature([]byte(passSalt),[]byte(userName),"sha512",true).(string)

	rt,_ := thisSql.Insert(
		"userLogin",
			[]string{"userID","userName","userPasswd","userSalt","userEmail","userPhone","userLevel","userCreateTime",
			"userLoginIP","lastLoginTime"},
			[]string{"0",userName,passHash,passSalt,email,"","1", lib.GetTimeDateTime(),c.ClientIP(), lib.GetTimeDateTime()},
		"",
		)

	if rt <= 0{
		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"Can not add new user"},
		)
		return
	}

	// Send auth email
	subject := appConfig.RegisterAuthData["subject"]
	urlApi := appConfig.RegisterAuthData["urlApi"]
	body := fmt.Sprintf(appConfig.RegisterAuthData["body"],userName,urlApi + activeApi)
	if lib.SendEmail(email,subject,body) != nil{
		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"Can not send email"},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		map[string]string{"status":"ok","contents":"Email has been send~"},
	)


}
// 邮箱激活码判断 POST
func api_userEnter_registerActive(c *gin.Context){

	activeCode := c.Param("activeCode")
	result := "激活码错误！"

	// SQL check
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	// Check unique
	r,_ := thisSql.Query(
		"userLogin",
		[]string{"userName","userSalt"},
		[]string{"active"},
		[]string{"0"},
		" AND ","",
	)
	for _,row := range r{
		activeHash := lib.GetHashSignature([]byte(row["userSalt"].(string)),[]byte(row["userName"].(string)),"sha512",
			true)

		if activeHash == activeCode{
			// Active account
			rt,_ := thisSql.Update(
				"userLogin",
				[]string{"active"},
				[]string{"1"},
				[]string{"userName","userSalt"},
				[]string{row["userName"].(string),row["userSalt"].(string)},
				" AND ",
				"",
			)
			if rt >= 0{
				result = "激活成功！"
				break
			}
		}
	}

	c.HTML(
		http.StatusOK,
		"userEnter_active.ni",
		gin.H{
			"content": result,
		},
	)

}

/*
	userForget
 */
// 用户密码找回 POST
func func_userEnter_findPass(c *gin.Context){

	step := c.DefaultPostForm("step","")

	switch step {
	case "1":
		authEmail := c.DefaultPostForm("authEmail","")

		// SQL find
		thisSql := &sqli.Sql{}
		thisSql.Connect()
		defer thisSql.Close()

		// Check unique
		r,_ := thisSql.Query(
			"userLogin",
			[]string{"*"},
			[]string{"userEmail"},
			[]string{authEmail},
			" AND ","",
		)

		if len(r) == 1{

			authCode := lib.GetHashSignature([]byte(lib.GetRandStr(32,"")),[]byte(authEmail),"sha256",true).(string)
			thisSql.Insert(
					"userForgetAuthCode",
					[]string{"email","activeCode","time"},
					[]string{authEmail,authCode,strconv.FormatInt(lib.GetTimeUnix(false) + 300,10) },
					"",
				)

			subject := appConfig.ForgetAuthData["subject"]
			body := fmt.Sprintf(appConfig.ForgetAuthData["body"],r[0]["userName"].(string),authCode)

			if lib.SendEmail(authEmail,subject,body) != nil{
				c.JSON(
					http.StatusOK,
					map[string]string{"status":"error","contents":"Can not send email"},
				)
				return
			}
			c.JSON(
				http.StatusOK,
				map[string]string{"status":"ok","contents":"Email has been send~"},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{"status":"ok","contents":"Email does not exist~"},
		)
	case "2":

		newPass := c.DefaultPostForm("newPass","")
		newRePass := c.DefaultPostForm("newRePass","")
		authCode := c.DefaultPostForm("authCode","")

		// cmp pass
		if newPass != newRePass{
			c.JSON(
				http.StatusOK,
				map[string]string{"status":"error","contents":"Different passwords"},
			)
			return
		}

		// SQL check
		thisSql := &sqli.Sql{}
		thisSql.Connect()
		defer thisSql.Close()

		// Check auth code
		r,_ := thisSql.Query(
			"userForgetAuthCode",
			[]string{"*"},
			[]string{"activeCode"},
			[]string{authCode},
			" AND ","",
		)

		if len(r) == 1{

			// Check time
			if r[0]["time"].(int64) < lib.GetTimeUnix(false){
				c.JSON(
					http.StatusOK,
					map[string]string{"status":"error","contents":"The email verification code has expired"},
				)
				return
			}

			// get user data and update
			passSalt := uuid.New()
			passHash := lib.GetHashSignature([]byte(newPass),[]byte(passSalt),"sha256",true).(string)

			rt,_ := thisSql.Update(
					"userLogin",
					[]string{"userPasswd","userSalt"},
					[]string{passHash,passSalt},
					[]string{"userEmail"},
					[]string{r[0]["email"].(string)},
					" AND ",
					"",
				)

			if rt <= 0{
				c.JSON(
					http.StatusOK,
					map[string]string{"status":"error","contents":"Failed to change password"},
				)
				return
			}
			c.JSON(
				http.StatusOK,
				map[string]string{"status":"ok","contents":"Password changed successfully~"},
				)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"The email verification code has expired"},
		)

	default:
		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"step error~"},
		)
	}
}