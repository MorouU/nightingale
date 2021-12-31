package appRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	appConfig "nightingale/config"
	lib "nightingale/lib/func"
	sqli "nightingale/lib/sql"
	"net/http"
	"strings"
)

// 管理员登录 GET
func pages_adminEnter_login(c *gin.Context){

	// SHOW
	if c.Request.Method == "GET" {

		if (!sessionKeyExist(sessionGet(c, "adminAuthCode"), "string")) || (sessionKeyExist(sessionGet(c, "adminAuthCodeTimeOut"), "int64") && lib.CmpAuthCodeTimeOut(sessionGet(c, "adminAuthCodeTimeOut").(int64), appConfig.LoginAuthCodeConfig["timeOut"].(int))) {
			authCode, _, timeOut := func_adminEnter_getLoginAuthCodeArray()
			sessionSet(c, map[interface{}]interface{}{
				"adminAuthCode":        authCode,
				"adminAuthCodeTimeOut": timeOut,
			})
		}

		authCodeB64 := func_adminEnter_getLoginAuthCodeB64(sessionGet(c, "adminAuthCode").(string))

		c.HTML(
			http.StatusOK,
			"adminEnter_login.ni",
			gin.H{
				"authCode": authCodeB64,
			},
		)
	}

	// LOGIN
	if c.Request.Method == "POST"{

		func_adminEnter_adminLogin(c)

	}
}

/*
	adminLogin
 */
// 验证码刷新
func api_adminEnter_loginAuthCodeRefresh(c *gin.Context){

	refresh := c.DefaultPostForm("refresh", "0")

	if refresh != "1"{
		c.JSON(
			http.StatusOK,
			map[string]string{"status":"error","contents":"params error ~"},
		)
	}

	authCode,authCodeB64,timeOut := func_adminEnter_getLoginAuthCodeArray()

	sessionSet(c, map[interface{}]interface{}{
		"adminAuthCode":authCode,
		"adminAuthCodeTimeOut": timeOut,
	})

	c.JSON(
		http.StatusOK,
		map[string]string{"status":"ok","contents":authCodeB64},
	)
}
// 获取验证码
func func_adminEnter_getLoginAuthCodeArray()(string,string,int64){

	length := appConfig.AdminAuthCodeConfig["length"].(int)
	chars := appConfig.AdminAuthCodeConfig["chars"].(string)
	width, height := appConfig.AdminAuthCodeConfig["width"].(int), appConfig.AdminAuthCodeConfig["height"].(int)
	timeOut := lib.GetTimeUnix(false) + int64(appConfig.AdminAuthCodeConfig["timeOut"].(int))
	authCode := lib.GetRandStr(length,chars)

	return authCode , lib.BaseEn(lib.GetAuthCode(width,height,authCode),64),timeOut
}
// 获取验证码图片base64
func func_adminEnter_getLoginAuthCodeB64(s string)string{
	return lib.BaseEn(lib.GetAuthCode(appConfig.AdminAuthCodeConfig["width"].(int), appConfig.AdminAuthCodeConfig["height"].(int),s),64)
}
// 管理员登录 POST
func func_adminEnter_adminLogin(c *gin.Context){

	adminAccount := c.DefaultPostForm("userAccount","")
	pass := c.DefaultPostForm("loginPass","")
	authCode := c.DefaultPostForm("authCode","")

	/*
		CHECK
	*/
	thisAuthCode := sessionGet(c,"adminAuthCode").(string)
	thisAuthCodeTime := sessionGet(c,"adminAuthCodeTimeOut").(int64)

	authCodeA,_,timeOut := func_adminEnter_getLoginAuthCodeArray()

	sessionSet(c, map[interface{}]interface{}{
		"adminAuthCode":authCodeA,
		"adminAuthCodeTimeOut": timeOut,
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
		"adminLogin",
		[]string{"*"},
		[]string{"adminName","adminEmail"},
		[]string{adminAccount,adminAccount},
		" OR ",
		"",
	)

	if len(r) == 1{
		passSalt := r[0]["adminSalt"].(string)
		passHash := lib.GetHashSignature([]byte(pass),[]byte(passSalt),"sha256",true).(string)


		if passHash != r[0]["adminPasswd"].(string){
			c.JSON(
				http.StatusOK,
				map[string]string{"status":"error","contents":"Incorrect username / email or password"},
			)
			return
		}

		// OK
		authID := uuid.New()
		sessionSet(c, map[interface{}]interface{}{
			"isAdmin":"1",
			"adminID":r[0]["adminID"].(int64),
			"adminName":r[0]["adminName"].(string),
			"adminEmail":r[0]["adminEmail"].(string),
			"authID":authID,
			"reqIP":c.ClientIP(),
		})

		// SQL Update
		rt,_ := thisSql.Update(
			"adminLogin",
			[]string{"lastLoginIP","lastLoginTime"},
			[]string{c.ClientIP(), lib.GetTimeDateTime()},
			[]string{"adminName","adminEmail"},
			[]string{r[0]["adminName"].(string),r[0]["adminEmail"].(string)},
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