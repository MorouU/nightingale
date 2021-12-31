package appRouter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	appConfig "nightingale/config"
	lib "nightingale/lib/func"
	sqli "nightingale/lib/sql"
	moduleLib "nightingale/module"
	"html"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// 管理员个人
func pages_adminProfile_profile(c *gin.Context){

	// CHECK LOGIN
	if !func_adminProfile_isLogin(c) || !func_adminProfile_checkAuthID(c){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/adminEnter/login",
		)
		return
	}

	// SHOW
	if c.Request.Method == "GET"{

		authID := sessionGet(c,"authID").(string)
		adminName := sessionGet(c,"adminName").(string)
		adminEmail := sessionGet(c,"adminEmail").(string)


		/*
			SQL
		*/
		// Connect
		thisSql := &sqli.Sql{}
		thisSql.Connect()
		defer thisSql.Close()

		// Get user data
		r,_ := thisSql.Query(
			"adminLogin",
			[]string{"adminPhone"},
			[]string{"adminName","adminEmail"},
			[]string{adminName,adminEmail},
			" AND ",
			"",
		)
		// ?
		if len(r) != 1{
			sessionClear(c)
			c.Redirect(
				http.StatusFound,
				"/adminEnter/login",
			)
		}

		adminPhone := r[0]["adminPhone"].(string)

		c.HTML(
			http.StatusOK,
			"adminProfile_adminedit.ni",
			gin.H{
				"authID": authID,
				"adminName":adminName,
				"adminEmail":adminEmail,
				"adminPhone":adminPhone,
			},
		)
	}

	// CHANGE
	if c.Request.Method == "POST"{

		func_adminProfile_changeInfo(c)

	}

}

// 管理员查看普通用户列表
func pages_adminProfile_userList(c *gin.Context){

	// CHECK LOGIN
	if !func_adminProfile_isLogin(c) || !func_adminProfile_checkAuthID(c){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/adminEnter/login",
		)
		return
	}

	// SHOW
	if c.Request.Method == "GET"{

		authID := sessionGet(c,"authID").(string)
		userRow := ""
		pagesTotal := 1
		userStatus := ""


		/*
			SQL
		*/
		// Connect
		thisSql := &sqli.Sql{}
		thisSql.Connect()
		defer thisSql.Close()

		// Get user data
		r,_ := thisSql.Query(
			"userLogin",
			[]string{"*"},
			[]string{},
			[]string{},
			" AND ",
			"",
		)
		total := len(r)
		showTotal := appConfig.UserListShowPerPage
		if total > 0 {
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.UserListShowPerPage)))
			if total < appConfig.UserListShowPerPage {
				showTotal = total
			}
			for _,row := range r[:showTotal]{
				if row["active"].(string) == "0"{
					userStatus = "停用"
				} else{
					userStatus = "启用"
				}
				userRow += fmt.Sprintf(moduleLib.UserListTableFormat,
					row["userID"].(string),
					html.EscapeString(row["userName"].(string)),
					html.EscapeString(row["userEmail"].(string)),
					userStatus,
					html.EscapeString(appConfig.UserLevelStringFromString[row["userLevel"].(string)]),
					html.EscapeString(row["userLoginIP"].(string)),
					html.EscapeString(row["lastLoginTime"].(string)),
					html.EscapeString(row["userCreateTime"].(string)),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
				)

			}

		}

		c.HTML(
			http.StatusOK,
			"adminProfile_userlist.ni",
			gin.H{
				"authID":authID,
				"userRow":template.HTML(userRow),
				"pagesTotal":pagesTotal,
			},
		)
	}

	// CHANGE
	if c.Request.Method == "POST"{

		func_adminProfile_controlUsers(c)

	}


}

// 管理员查看普通用户任务
func pages_adminProfile_userTasks(c *gin.Context){

	// CHECK LOGIN
	if !func_adminProfile_isLogin(c) || !func_adminProfile_checkAuthID(c){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/adminEnter/login",
		)
		return
	}

	// SHOW
	if c.Request.Method == "GET"{

		authID := sessionGet(c,"authID").(string)
		adminID := strconv.FormatInt(sessionGet(c,"adminID").(int64),10)
		userID := c.Param("userID")
		taskRow := ""
		taskStatus := ""
		pagesTotal := 1


		/*
			SQL
		*/
		// Connect
		thisSql := &sqli.Sql{}
		thisSql.Connect()
		defer thisSql.Close()

		// get total
		r,_ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskUserID"},
			[]string{userID},
			" AND ",
			"",
		)

		total := len(r)
		showTotal := appConfig.TaskListShowPerPage
		if total > 0 {

			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.TaskListShowPerPage)))
			if total < appConfig.TaskListShowPerPage {
				showTotal = total
			}

			for _,row := range r[:showTotal]{

				if row["taskStatus"].(int64) != 1{
					taskStatus = "停用"
				}else{
					taskStatus = "启用"
				}

				taskRow += fmt.Sprintf(moduleLib.TaskListTableFormat,
					row["taskID"].(int64),
					html.EscapeString(row["taskName"].(string)),
					html.EscapeString(row["taskModuleName"].(string)),
					taskStatus,
					html.EscapeString(row["taskCreateTime"].(string)),
					row["taskID"].(int64),
					row["taskID"].(int64),
					row["taskID"].(int64),
					row["taskID"].(int64),

				)

			}

		}

		// get user data

		rt,_ := thisSql.Query(
			"userLogin",
			[]string{"*"},
			[]string{"userID"},
			[]string{userID},
			" AND ",
			"",
		)

		if len(rt) != 1{

			c.Redirect(
				http.StatusFound,
				"/adminEnter/"+adminID+"/userList",
			)

		}

		c.HTML(
			http.StatusOK,
			"adminProfile_tasklist.ni",
			gin.H{
				"authID":authID,
				"userID":userID,
				"userName":rt[0]["userName"].(string),
				"taskRow":template.HTML(taskRow),
				"pagesTotal":pagesTotal,
			},
		)

	}

	// VIEW
	if c.Request.Method == "POST"{

		func_adminProfile_controlTasks(c)

	}



}

// 管理员查看普通用户模块
func pages_adminProfile_userModules(c *gin.Context){

	// CHECK LOGIN
	if !func_adminProfile_isLogin(c) || !func_adminProfile_checkAuthID(c){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/adminEnter/login",
		)
		return
	}

	// SHOW
	if c.Request.Method == "GET"{

		authID := sessionGet(c,"authID").(string)
		adminID := strconv.FormatInt(sessionGet(c,"adminID").(int64),10)
		userID := c.Param("userID")
		moduleRow := ""
		pagesTotal := 1


		/*
			SQL
		*/
		// Connect
		thisSql := &sqli.Sql{}
		thisSql.Connect()
		defer thisSql.Close()

		// get total
		r,_ := thisSql.Query(
			"userModule",
			[]string{"*"},
			[]string{"belongUserID"},
			[]string{userID},
			" AND ",
			"",
		)

		total := len(r)
		showTotal := appConfig.ModuleListShowPerPage
		if total > 0 {

			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.ModuleListShowPerPage)))
			if total < appConfig.ModuleListShowPerPage {
				showTotal = total
			}

			for _,row := range r[:showTotal]{


				moduleRow += fmt.Sprintf(moduleLib.ModuleListTableFormat,
					row["moduleID"].(int64),
					html.EscapeString(row["moduleName"].(string)),
					html.EscapeString(row["createTime"].(string)),
					row["moduleID"].(int64),
					row["moduleID"].(int64),
					row["moduleID"].(int64),
					row["moduleID"].(int64),

				)

			}

		}

		// get user data

		rt,_ := thisSql.Query(
			"userLogin",
			[]string{"*"},
			[]string{"userID"},
			[]string{userID},
			" AND ",
			"",
		)

		if len(rt) != 1{

			c.Redirect(
				http.StatusFound,
				"/adminEnter/"+adminID+"/userList",
			)

		}

		c.HTML(
			http.StatusOK,
			"adminProfile_modulelist.ni",
			gin.H{
				"authID":authID,
				"userID":userID,
				"userName":rt[0]["userName"].(string),
				"taskRow":template.HTML(moduleRow),
				"pagesTotal":pagesTotal,
			},
		)

	}

	// VIEW
	if c.Request.Method == "POST"{

		func_adminProfile_controlModules(c)

	}

}

// 管理员登出
func pages_adminProfile_logout(c *gin.Context){

	// CHECK LOGIN
	if !func_adminProfile_isLogin(c) || !func_adminProfile_checkAuthID(c){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/adminEnter/login",
		)
		return
	}

	sessionClear(c)
	c.Redirect(
		http.StatusFound,
		"/adminEnter/login",
	)

}
/*
	adminProfile
 */
// 判断管理员是否登录
func func_adminProfile_isLogin(c *gin.Context) bool {

	return sessionKeyExist(sessionGet(c,"isAdmin"),"string") && sessionGet(c,"isAdmin").(string) == "1"

}
// 判断管理员认证ID是否合法
func func_adminProfile_checkAuthID(c *gin.Context) bool{

	return sessionKeyExist(sessionGet(c,"authID"),"string") && c.Param("uuid") == sessionGet(c,"authID").(string)

}
// 判断请求参数是否和认证ID一致
func func_adminProfile_checkAuthIDFromRequest(c *gin.Context,method,key string) bool{

	switch method {
	case	"GET":
		return sessionKeyExist(sessionGet(c,"authID"),"string") && c.DefaultQuery(key,"NoKey") == sessionGet(c,"authID").(string)
	default:
		return sessionKeyExist(sessionGet(c,"authID"),"string") && c.DefaultPostForm(key,"NoKey") == sessionGet(c,"authID").(string)
	}
}

// 管理员信息修改
func func_adminProfile_changeInfo(c *gin.Context){

	// CHECK LOGIN
	if !func_adminProfile_isLogin(c) || !func_adminProfile_checkAuthID(c) || !func_adminProfile_checkAuthIDFromRequest(c,"POST","authID"){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/adminEnter/login",
		)
		return
	}

	// check param
	step := c.DefaultPostForm("step","")
	if step != "info" && step != "pass"{
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"error",
				"contents":"step error",
			},
		)
		return
	}

	adminName := sessionGet(c,"adminName").(string)
	adminEmail := sessionGet(c,"adminEmail").(string)

	/*
		SQL
	*/
	// Connect
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	switch step {

	case "info":
		editAdminName := c.DefaultPostForm("editAdminName","")
		editAdminEmail := c.DefaultPostForm("editAdminEmail","")
		editAdminPhone := c.DefaultPostForm("editAdminPhone","")


		// Check Username
		if !lib.CheckUsername(editAdminName){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"The user name is incorrectly formatted",
				},
			)
			return
		}

		// Check Email
		if !lib.VerifyEmailFormat(editAdminEmail){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Incorrect email format",
				},
			)
			return
		}

		// Check Phone
		if !lib.VerifyMobileFormat(editAdminPhone){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"The format of mobile phone number is wrong",
				},
			)
			return
		}

		// Update

		r,_:=thisSql.Update(
			"adminLogin",
			[]string{"adminName","adminEmail","adminPhone"},
			[]string{editAdminName,editAdminEmail,editAdminPhone},
			[]string{"adminName","adminEmail"},
			[]string{adminName,adminEmail},
			" AND ",
			"",
		)

		if r == 0{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"warning",
					"contents":"The information may not have changed",
					"adminName":editAdminName,
					"adminEmail":editAdminEmail,
					"adminPhone":editAdminPhone,
				},
			)
			return
		}
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"contents":"Successful change~",
				"adminName":editAdminName,
				"adminEmail":editAdminEmail,
				"adminPhone":editAdminPhone,
			},
		)

	case "pass":

		changeAdminCPass := c.DefaultPostForm("changeAdminCPass","")
		changeAdminNPass := c.DefaultPostForm("changeAdminNPass","")
		changeAdminVPass := c.DefaultPostForm("changeAdminVPass","")

		// Check New pass
		if changeAdminNPass != changeAdminVPass{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Different passwords",
				},
			)
			return
		}

		// Check Current pass
		r,_ := thisSql.Query(
			"adminLogin",
			[]string{"adminPasswd","adminSalt"},
			[]string{"adminName","adminEmail"},
			[]string{adminName,adminEmail},
			" AND ",
			"",
		)
		// ?
		if len(r) != 1{
			sessionClear(c)
			c.Redirect(
				http.StatusFound,
				"/adminEnter/login",
			)
		}
		cPassHash := lib.GetHashSignature([]byte(changeAdminCPass),[]byte(r[0]["adminSalt"].(string)),"sha256",true).(string)
		// Cmp Pass hash
		if cPassHash != r[0]["adminPasswd"].(string){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Incorrect password",
				},
			)
			return
		}

		// update
		nSalt := uuid.New()
		nPassHash := lib.GetHashSignature([]byte(changeAdminNPass),[]byte(nSalt),"sha256",true).(string)

		rt,_ := thisSql.Update(
			"adminLogin",
			[]string{"adminPasswd","adminSalt"},
			[]string{nPassHash,nSalt},
			[]string{"adminName","adminEmail"},
			[]string{adminName,adminEmail},
			" AND ",
			"",
		)

		if rt == 0{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"warning",
					"contents":"The password may not have changed",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"contents":"Password changed successfully~",
			},
		)
	}

}

/*
	adminUserList
 */
// 管理员管理用户操作
func func_adminProfile_controlUsers(c *gin.Context){

	// CHECK LOGIN
	if !func_adminProfile_isLogin(c) || !func_adminProfile_checkAuthID(c) || !func_adminProfile_checkAuthIDFromRequest(c,"POST","authID"){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/adminEnter/login",
		)
		return
	}

	// check param
	step := c.DefaultPostForm("step","")

	if step != "previous" && step != "next" && step != "back" && step != "find" && step != "control" && step != "delete" && step != "view" && step != "new" && step != "edit"{
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"error",
				"contents":"step error",
			},
		)
		return
	}

	/*
		SQL
	*/
	// Connect
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	switch step {
	case "previous":

		currentPage,_ := strconv.Atoi(c.DefaultPostForm("currentPage",""))
		pagesTotal := 0
		userRow := ""
		userStatus := ""

		// get total
		r,_ := thisSql.Query(
			"userLogin",
			[]string{"*"},
			[]string{},
			[]string{},
			" AND ",
			"",
		)

		total := len(r)
		if total > 0 {

			endPages := (currentPage - 1) * appConfig.UserListShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.UserListShowPerPage)))

			if currentPage - 1 > 0 && currentPage - 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage - 2)*appConfig.UserListShowPerPage :endPages]{


					if row["active"].(string) == "0"{
						userStatus = "停用"
					} else{
						userStatus = "启用"
					}



					userRow += fmt.Sprintf(moduleLib.UserListTableFormat,
						row["userID"].(string),
						html.EscapeString(row["userName"].(string)),
						html.EscapeString(row["userEmail"].(string)),
						userStatus,
						html.EscapeString(appConfig.UserLevelStringFromString[row["userLevel"].(string)]),
						html.EscapeString(row["userLoginIP"].(string)),
						html.EscapeString(row["lastLoginTime"].(string)),
						html.EscapeString(row["userCreateTime"].(string)),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
					)

				}


				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"html":userRow,
						"page":strconv.Itoa(currentPage - 1),
						"total":strconv.Itoa(pagesTotal),
					},
				)

			}

		}



	case "next":

		currentPage,_ := strconv.Atoi(c.DefaultPostForm("currentPage",""))
		pagesTotal := 0
		userRow := ""
		userStatus := ""

		// get total
		r,_ := thisSql.Query(
			"userLogin",
			[]string{"*"},
			[]string{},
			[]string{},
			" AND ",
			"",
		)

		total := len(r)
		if total > 0 {

			endPages := (currentPage + 1) * appConfig.UserListShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.UserListShowPerPage)))

			if currentPage + 1 > 1 && currentPage + 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage)*appConfig.UserListShowPerPage :endPages]{


					if row["active"].(string) == "0"{
						userStatus = "停用"
					} else{
						userStatus = "启用"
					}


					userRow += fmt.Sprintf(moduleLib.UserListTableFormat,
						row["userID"].(string),
						html.EscapeString(row["userName"].(string)),
						html.EscapeString(row["userEmail"].(string)),
						userStatus,
						html.EscapeString(appConfig.UserLevelStringFromString[row["userLevel"].(string)]),
						html.EscapeString(row["userLoginIP"].(string)),
						html.EscapeString(row["lastLoginTime"].(string)),
						html.EscapeString(row["userCreateTime"].(string)),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
						row["userID"].(string),
					)

				}


				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"html":userRow,
						"page":strconv.Itoa(currentPage + 1),
						"total":strconv.Itoa(pagesTotal),
					},
				)

			}

		}




	case "back":

		userRow := ""
		pagesTotal := 1
		userStatus := ""

		// get total
		r,_ := thisSql.Query(
			"userLogin",
			[]string{"*"},
			[]string{},
			[]string{},
			" AND ",
			"",
		)

		total := len(r)
		showTotal := appConfig.UserListShowPerPage
		if total > 0 {

			pagesTotal = int(math.Ceil(float64(total) / float64(appConfig.UserListShowPerPage)))
			if total < appConfig.UserListShowPerPage {
				showTotal = total
			}

			for _, row := range r[:showTotal] {

				if row["active"].(string) == "0"{
					userStatus = "停用"
				} else{
					userStatus = "启用"
				}

				userRow += fmt.Sprintf(moduleLib.UserListTableFormat,
					row["userID"].(string),
					html.EscapeString(row["userName"].(string)),
					html.EscapeString(row["userEmail"].(string)),
					userStatus,
					html.EscapeString(appConfig.UserLevelStringFromString[row["userLevel"].(string)]),
					html.EscapeString(row["userLoginIP"].(string)),
					html.EscapeString(row["lastLoginTime"].(string)),
					html.EscapeString(row["userCreateTime"].(string)),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
					row["userID"].(string),
				)
			}
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"html":userRow,
				"page": strconv.Itoa(1),
				"total":strconv.Itoa(pagesTotal),
			},
		)


	case "find":

		account := c.DefaultPostForm("findAccount","")
		userRow := ""
		userStatus := ""


		// get total
		r,_ := thisSql.Query(
			"userLogin",
			[]string{"*"},
			[]string{"userID","userName","userEmail"},
			[]string{account,account,account},
			" OR ",
			"",
		)


		if len(r) > 0{

			for _,row := range r{

				if row["active"].(int64) == 0{
					userStatus = "停用"
				} else{
					userStatus = "启用"
				}

				userRow += fmt.Sprintf(moduleLib.UserListTableFormat,
					strconv.FormatInt(row["userID"].(int64),10),
					html.EscapeString(row["userName"].(string)),
					html.EscapeString(row["userEmail"].(string)),
					userStatus,
					html.EscapeString(appConfig.UserLevelString[row["userLevel"].(int64)]),
					html.EscapeString(row["userLoginIP"].(string)),
					html.EscapeString(row["lastLoginTime"].(string)),
					html.EscapeString(row["userCreateTime"].(string)),
					strconv.FormatInt(row["userID"].(int64),10),
					strconv.FormatInt(row["userID"].(int64),10),
					strconv.FormatInt(row["userID"].(int64),10),
					strconv.FormatInt(row["userID"].(int64),10),
					strconv.FormatInt(row["userID"].(int64),10),
					strconv.FormatInt(row["userID"].(int64),10),
					strconv.FormatInt(row["userID"].(int64),10),
					strconv.FormatInt(row["userID"].(int64),10),
				)
			}


			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"ok",
					"html":userRow,
					"page": strconv.Itoa(1),
					"total":strconv.Itoa(1),
				},
			)

		}else{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"ok",
					"html":"",
					"page": strconv.Itoa(1),
					"total":strconv.Itoa(1),
				},
			)

		}

	case "view":

		userID := c.DefaultPostForm("userID","")
		userStatus := "1"

		// Get status
		r,_ := thisSql.Query(
			"userLogin",
			[]string{"*"},
			[]string{"userID"},
			[]string{userID},
			" AND ",
			"",
		)

		if len(r) != 1{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"User ID error",
				},
			)
			return

		}

		if r[0]["active"].(int64) != 1{
			userStatus = "0"
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"userName":r[0]["userName"].(string),
				"userEmail":r[0]["userEmail"].(string),
				"userPasswd":"",
				"userStatus":userStatus,
				"userLevel":strconv.FormatInt(r[0]["userLevel"].(int64),10),
			},
		)


	case "delete":
		userID := c.DefaultPostForm("userID","")
		// delete
		r,_ := thisSql.Delete(
			"userLogin",
			[]string{"userID"},
			[]string{userID},
			" AND ",
			"",
		)
		if r !=1 {

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Delete user failed",
				},
			)
			return
		}
		thisSql.Delete(
			"appTaskView",
			[]string{"userID"},
			[]string{userID},
			" AND ",
			"",
		)
		thisSql.Delete(
			"appTasks",
			[]string{"taskUserID"},
			[]string{userID},
			" AND ",
			"",
		)
		thisSql.Delete(
			"userModule",
			[]string{"belongUserID"},
			[]string{userID},
			" AND ",
			"",
		)
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"Delete user successfully~",
			},
		)

	case "new":

		userName := c.DefaultPostForm("userName","")
		userEmail := c.DefaultPostForm("userEmail","")
		userPhone := c.DefaultPostForm("userPhone","")
		userPass := c.DefaultPostForm("userPass","")
		userStatus := c.DefaultPostForm("userStatus","")
		userLevel := c.DefaultPostForm("userLevel","")

		// check

		if ! lib.CheckUsername(userName){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"The user name is not valid",
				},
			)
			return
		}

		// Check Email
		if !lib.VerifyEmailFormat(userEmail){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Incorrect email format",
				},
			)
			return
		}

		// Check Phone
		if len(userPhone) > 0 &&!lib.VerifyMobileFormat(userPhone){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"The format of mobile phone number is wrong",
				},
			)
			return
		}
		// Check Status
		if userStatus != "1"{
			userStatus = "0"
		}
		// Check Level
		if userLevel != "0" && userLevel != "1" && userLevel != "2"{
			userLevel = "0"
		}
		// Check unique
		if func_userEnter_checkUnique(userName,userEmail){
			c.JSON(
				http.StatusOK,
				map[string]string{"status":"error","contents":"Username/email already exists"},
			)
			return
		}
		// Get user data and insert
		passSalt := uuid.New()
		passHash := lib.GetHashSignature([]byte(userPass),[]byte(passSalt),"sha256",true).(string)
		r,_ := thisSql.Insert(
			"userLogin",
				[]string{"userID","userName","userPasswd","userSalt","userEmail","userPhone","userLevel","userCreateTime",
				"userLoginIP","lastLoginTime","active"},
				[]string{"0",userName,passHash,passSalt,userEmail,userPhone,userLevel, lib.GetTimeDateTime(),c.ClientIP(),
					lib.GetTimeDateTime(),userStatus},
			"",
		)

		if r <= 0{
			c.JSON(
				http.StatusOK,
				map[string]string{"status":"error","contents":"Can not add new user"},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{"status":"refresh","contents":"User created successfully~"},
		)

	case "edit":

		userID := c.DefaultPostForm("userID","")
		userName := c.DefaultPostForm("userName","")
		userEmail := c.DefaultPostForm("userEmail","")
		userPhone := c.DefaultPostForm("userPhone","")
		userPass := c.DefaultPostForm("userPass","")
		userStatus := c.DefaultPostForm("userStatus","")
		userLevel := c.DefaultPostForm("userLevel","")

		// Check Username
		if !lib.CheckUsername(userName){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"The user name is incorrectly formatted",
				},
			)
			return
		}

		// Check Email
		if !lib.VerifyEmailFormat(userEmail){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Incorrect email format",
				},
			)
			return
		}

		// Check Phone
		if len(userPhone) > 0 &&!lib.VerifyMobileFormat(userPhone){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"The format of mobile phone number is wrong",
				},
			)
			return
		}

		// Check Status
		if userStatus != "1"{
			userStatus = "0"
		}

		// Check Level
		if userLevel != "0" && userLevel != "1" && userLevel != "2"{
			userLevel = "0"
		}

		// Check Pass
		if len(userPass) > 0{

			passSalt := uuid.New()
			passHash := lib.GetHashSignature([]byte(userPass),[]byte(passSalt),"sha256",true).(string)

			r,_:=thisSql.Update(
					"userLogin",
					[]string{"userName","userEmail","userPasswd","userSalt","userPhone","active","userLevel"},
					[]string{userName,userEmail,passHash,passSalt,userPhone,userStatus,userLevel},
					[]string{"userID"},
					[]string{userID},
					" AND ",
					"",
				)
			// ?
			if r != 1{
				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"warning",
						"contents":"User information may not have changed",
					},
				)
			}

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"refresh",
					"contents":"Save user successfully~",
				},
			)

		}else{

			r,_:=thisSql.Update(
				"userLogin",
				[]string{"userName","userEmail","userPhone","active","userLevel"},
				[]string{userName,userEmail,userPhone,userStatus,userLevel},
				[]string{"userID"},
				[]string{userID},
				" AND ",
				"",
			)
			// ?
			if r != 1{
				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"warning",
						"contents":"User information may not have changed",
					},
				)
				return
			}

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"refresh",
					"contents":"Save user successfully~",
				},
			)

		}

	}

}

/*
	adminUserTaskList
 */
// 管理员管理用户任务操作
func func_adminProfile_controlTasks(c *gin.Context){

	// CHECK LOGIN
	if !func_adminProfile_isLogin(c) || !func_adminProfile_checkAuthID(c) || !func_adminProfile_checkAuthIDFromRequest(c,"POST","authID"){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/adminEnter/login",
		)
		return
	}

	userID := c.Param("userID")

	// check param
	step := c.DefaultPostForm("step","")

	if step != "previous" && step != "next" && step != "back" && step != "find"  && step != "delete" && step != "view" && step != "save" {
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"error",
				"contents":"step error",
			},
		)
		return
	}

	/*
		SQL
	*/
	// Connect
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	switch step {
	case "previous":

		currentPage,_ := strconv.Atoi(c.DefaultPostForm("currentPage",""))
		pagesTotal := 0
		taskRow := ""
		taskStatus := ""

		// get total
		r,_ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskUserID"},
			[]string{userID},
			" AND ",
			"",
		)

		total := len(r)
		if total > 0 {

			endPages := (currentPage - 1) * appConfig.TaskListShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.TaskListShowPerPage)))

			if currentPage - 1 > 0 && currentPage - 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage - 2)*appConfig.TaskListShowPerPage :endPages]{


					if row["taskStatus"].(int64) == 0{
						taskStatus = "停用"
					} else{
						taskStatus = "启用"
					}



					taskRow += fmt.Sprintf(moduleLib.TaskListTableFormat,
						row["taskID"].(int64),
						html.EscapeString(row["taskName"].(string)),
						html.EscapeString(row["taskModuleName"].(string)),
						taskStatus,
						html.EscapeString(row["taskCreateTime"].(string)),
						row["taskID"].(int64),
						row["taskID"].(int64),
						row["taskID"].(int64),
						row["taskID"].(int64),
					)

				}


				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"html":taskRow,
						"page":strconv.Itoa(currentPage - 1),
						"total":strconv.Itoa(pagesTotal),
					},
				)

			}

		}



	case "next":

		currentPage,_ := strconv.Atoi(c.DefaultPostForm("currentPage",""))
		pagesTotal := 0
		taskRow := ""
		taskStatus := ""

		// get total
		r,_ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskUserID"},
			[]string{userID},
			" AND ",
			"",
		)

		total := len(r)
		if total > 0 {

			endPages := (currentPage + 1) * appConfig.TaskListShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.TaskListShowPerPage)))

			if currentPage + 1 > 1 && currentPage + 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage)*appConfig.TaskListShowPerPage :endPages]{


					if row["taskStatus"].(int64) == 0{
						taskStatus = "停用"
					} else{
						taskStatus = "启用"
					}


					taskRow += fmt.Sprintf(moduleLib.TaskListTableFormat,
						row["taskID"].(int64),
						html.EscapeString(row["taskName"].(string)),
						html.EscapeString(row["taskModuleName"].(string)),
						taskStatus,
						html.EscapeString(row["taskCreateTime"].(string)),
						row["taskID"].(int64),
						row["taskID"].(int64),
						row["taskID"].(int64),
						row["taskID"].(int64),
					)

				}


				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"html":taskRow,
						"page":strconv.Itoa(currentPage + 1),
						"total":strconv.Itoa(pagesTotal),
					},
				)

			}

		}




	case "back":

		taskRow := ""
		pagesTotal := 1
		taskStatus := ""

		// get total
		r,_ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskUserID"},
			[]string{userID},
			" AND ",
			"",
		)

		total := len(r)
		showTotal := appConfig.TaskListShowPerPage
		if total > 0 {

			pagesTotal = int(math.Ceil(float64(total) / float64(appConfig.TaskListShowPerPage)))
			if total < appConfig.TaskListShowPerPage {
				showTotal = total
			}

			for _, row := range r[:showTotal] {

				if row["taskStatus"].(int64) == 0{
					taskStatus = "停用"
				} else{
					taskStatus = "启用"
				}

				taskRow += fmt.Sprintf(moduleLib.TaskListTableFormat,
					row["taskID"].(int64),
					html.EscapeString(row["taskName"].(string)),
					html.EscapeString(row["taskModuleName"].(string)),
					taskStatus,
					html.EscapeString(row["taskCreateTime"].(string)),
					row["taskID"].(int64),
					row["taskID"].(int64),
					row["taskID"].(int64),
					row["taskID"].(int64),
				)
			}
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"html":taskRow,
				"page": strconv.Itoa(1),
				"total":strconv.Itoa(pagesTotal),
			},
		)


	case "find":

		task := c.DefaultPostForm("findTask","")
		taskRow := ""
		taskStatus := ""


		// get total
		r,_ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskUserID"},
			[]string{userID},
			" OR ",
			"",
		)


		if len(r) > 0{

			for _,row := range r{

				if strconv.FormatInt(row["taskID"].(int64),10) == task || row["taskName"].(string) == task {

					if row["taskStatus"].(int64) == 0 {
						taskStatus = "停用"
					} else {
						taskStatus = "启用"
					}

					taskRow += fmt.Sprintf(moduleLib.TaskListTableFormat,
						row["taskID"].(int64),
						html.EscapeString(row["taskName"].(string)),
						html.EscapeString(row["taskModuleName"].(string)),
						taskStatus,
						html.EscapeString(row["taskCreateTime"].(string)),
						row["taskID"].(int64),
						row["taskID"].(int64),
						row["taskID"].(int64),
						row["taskID"].(int64),
					)
				}
			}


			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"ok",
					"html":taskRow,
					"page": strconv.Itoa(1),
					"total":strconv.Itoa(1),
				},
			)

		}else{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"ok",
					"html":"",
					"page": strconv.Itoa(1),
					"total":strconv.Itoa(1),
				},
			)

		}

	case "view":

		taskID := c.DefaultPostForm("taskID","")
		taskStatus := "1"

		// Get status
		r,_ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskUserID","taskID"},
			[]string{userID,taskID},
			" AND ",
			"",
		)

		if len(r) != 1{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Task ID / name or user ID error",
				},
			)
			return

		}

		if r[0]["taskStatus"].(int64) != 1{
			taskStatus = "0"
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"taskName":r[0]["taskName"].(string),
				"taskData":r[0]["taskData"].(string),
				"taskStatus":taskStatus,
			},
		)


	case "delete":

		taskID := c.DefaultPostForm("taskID","")

		// delete
		r,_ := thisSql.Delete(
			"appTasks",
			[]string{"taskID","taskUserID"},
			[]string{taskID,userID},
			" AND ",
			"",
		)

		if r !=1 {

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Delete task failed",
				},
			)
			return

		}

		thisSql.Delete(
			"appTaskView",
			[]string{"taskID","userID"},
			[]string{taskID,userID},
			" AND ",
			"",
		)


		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"Delete task successfully~",
			},
		)

	case "save":

		taskID := c.DefaultPostForm("taskID","")
		taskName := c.DefaultPostForm("taskName","")
		taskData := c.DefaultPostForm("taskData","")
		taskStatus := c.DefaultPostForm("taskStatus","")


		// check

		if ! lib.CheckTaskName(taskName){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"The task name is not valid",
				},
			)
			return
		}

		// Check Status
		if taskStatus != "1"{
			taskStatus = "0"
		}


		// update
		r,_ := thisSql.Update(
				"appTasks",
				[]string{"taskName","taskData","taskStatus"},
				[]string{taskName,taskData,taskStatus},
				[]string{"taskID","taskUserID"},
				[]string{taskID,userID},
				" AND ",
				"",
			)

		if r <= 0{
			c.JSON(
				http.StatusOK,
				map[string]string{"status":"warning","contents":"The task information may not have changed"},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{"status":"refresh","contents":"task modify successfully~"},
		)



	}

}

/*
	adminUserModuleList
 */
// 管理员管理用户模块操作
func func_adminProfile_controlModules(c *gin.Context){

	// CHECK LOGIN
	if !func_adminProfile_isLogin(c) || !func_adminProfile_checkAuthID(c) || !func_adminProfile_checkAuthIDFromRequest(c,"POST","authID"){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/adminEnter/login",
		)
		return
	}

	userID := c.Param("userID")

	// check param
	step := c.DefaultPostForm("step","")

	if step != "previous" && step != "next" && step != "back" && step != "find"  && step != "delete" && step != "view" && step != "save" {
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"error",
				"contents":"step error",
			},
		)
		return
	}

	/*
		SQL
	*/
	// Connect
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	switch step {
	case "previous":

		currentPage,_ := strconv.Atoi(c.DefaultPostForm("currentPage",""))
		pagesTotal := 0
		moduleRow := ""

		// get total
		r,_ := thisSql.Query(
			"userModule",
			[]string{"*"},
			[]string{"belongUserID"},
			[]string{userID},
			" AND ",
			"",
		)

		total := len(r)
		if total > 0 {

			endPages := (currentPage - 1) * appConfig.ModuleListShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.ModuleListShowPerPage)))

			if currentPage - 1 > 0 && currentPage - 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage - 2)*appConfig.ModuleListShowPerPage :endPages]{



					moduleRow += fmt.Sprintf(moduleLib.ModuleListTableFormat,
						row["moduleID"].(int64),
						html.EscapeString(row["moduleName"].(string)),
						html.EscapeString(row["createTime"].(string)),
						row["moduleID"].(int64),
						row["moduleID"].(int64),
						row["moduleID"].(int64),
						row["moduleID"].(int64),
					)

				}


				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"html":moduleRow,
						"page":strconv.Itoa(currentPage - 1),
						"total":strconv.Itoa(pagesTotal),
					},
				)

			}

		}



	case "next":

		currentPage,_ := strconv.Atoi(c.DefaultPostForm("currentPage",""))
		pagesTotal := 0
		moduleRow := ""

		// get total
		r,_ := thisSql.Query(
			"userModule",
			[]string{"*"},
			[]string{"belongUserID"},
			[]string{userID},
			" AND ",
			"",
		)

		total := len(r)
		if total > 0 {

			endPages := (currentPage + 1) * appConfig.ModuleListShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.ModuleListShowPerPage)))

			if currentPage + 1 > 1 && currentPage + 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage)*appConfig.ModuleListShowPerPage :endPages]{



					moduleRow += fmt.Sprintf(moduleLib.ModuleListTableFormat,
						row["moduleID"].(int64),
						html.EscapeString(row["moduleName"].(string)),
						html.EscapeString(row["createTime"].(string)),
						row["moduleID"].(int64),
						row["moduleID"].(int64),
						row["moduleID"].(int64),
						row["moduleID"].(int64),
					)

				}


				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"html":moduleRow,
						"page":strconv.Itoa(currentPage + 1),
						"total":strconv.Itoa(pagesTotal),
					},
				)

			}

		}




	case "back":

		moduleRow := ""
		pagesTotal := 1

		// get total
		r,_ := thisSql.Query(
			"userModule",
			[]string{"*"},
			[]string{"belongUserID"},
			[]string{userID},
			" AND ",
			"",
		)

		total := len(r)
		showTotal := appConfig.ModuleListShowPerPage
		if total > 0 {

			pagesTotal = int(math.Ceil(float64(total) / float64(appConfig.ModuleListShowPerPage)))
			if total < appConfig.ModuleListShowPerPage {
				showTotal = total
			}

			for _, row := range r[:showTotal] {

				moduleRow += fmt.Sprintf(moduleLib.ModuleListTableFormat,
					row["moduleID"].(int64),
					html.EscapeString(row["moduleName"].(string)),
					html.EscapeString(row["createTime"].(string)),
					row["moduleID"].(int64),
					row["moduleID"].(int64),
					row["moduleID"].(int64),
					row["moduleID"].(int64),
				)
			}
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"html":moduleRow,
				"page": strconv.Itoa(1),
				"total":strconv.Itoa(pagesTotal),
			},
		)


	case "find":

		module := c.DefaultPostForm("findModule","")
		moduleRow := ""


		// get total
		r,_ := thisSql.Query(
			"userModule",
			[]string{"*"},
			[]string{"belongUserID"},
			[]string{userID},
			" AND ",
			"",
		)


		if len(r) > 0{

			for _,row := range r{

				if strconv.FormatInt(row["moduleID"].(int64),10) == module || row["moduleName"].(string) == module {

					moduleRow += fmt.Sprintf(moduleLib.ModuleListTableFormat,
						row["moduleID"].(int64),
						html.EscapeString(row["moduleName"].(string)),
						html.EscapeString(row["createTime"].(string)),
						row["moduleID"].(int64),
						row["moduleID"].(int64),
						row["moduleID"].(int64),
						row["moduleID"].(int64),
					)
				}
			}


			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"ok",
					"html":moduleRow,
					"page": strconv.Itoa(1),
					"total":strconv.Itoa(1),
				},
			)

		}else{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"ok",
					"html":"",
					"page": strconv.Itoa(1),
					"total":strconv.Itoa(1),
				},
			)

		}

	case "view":

		moduleID := c.DefaultPostForm("moduleID","")

		// Get status
		r,_ := thisSql.Query(
			"userModule",
			[]string{"*"},
			[]string{"moduleID","belongUserID"},
			[]string{moduleID,userID},
			" AND ",
			"",
		)

		if len(r) != 1{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Module ID / name or user ID error",
				},
			)
			return

		}


		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"moduleName":r[0]["moduleName"].(string),
				"moduleData":r[0]["moduleData"].(string),
				"moduleCode":r[0]["moduleCode"].(string),
			},
		)


	case "delete":

		moduleID := c.DefaultPostForm("moduleID","")

		// delete
		r,_ := thisSql.Delete(
			"userModule",
			[]string{"moduleID","belongUserID"},
			[]string{moduleID,userID},
			" AND ",
			"",
		)

		if r !=1 {

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Delete module failed",
				},
			)
			return

		}

		thisSql.Delete(
			"appTasks",
			[]string{"taskModuleID","userID","taskModulePublic"},
			[]string{moduleID,userID,"0"},
			" AND ",
			"",
		)

		thisSql.Delete(
			"appTaskView",
			[]string{"moduleID","userID","modulePublic"},
			[]string{moduleID,userID,"0"},
			" AND ",
			"",
		)


		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"Delete module successfully~",
			},
		)

	case "save":

		moduleID := c.DefaultPostForm("moduleID","")
		moduleName := c.DefaultPostForm("moduleName","")
		moduleData := c.DefaultPostForm("moduleData","")
		moduleCode := c.DefaultPostForm("moduleCode","")


		// check

		if ! lib.CheckCustomModuleName(moduleName){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"The module name is not valid",
				},
			)
			return
		}


		// update
		r,_ := thisSql.Update(
			"userModule",
			[]string{"moduleName","moduleData","moduleCode"},
			[]string{moduleName,moduleData,moduleCode},
			[]string{"moduleID","belongUserID"},
			[]string{moduleID,userID},
			" AND ",
			"",
		)

		if r <= 0{
			c.JSON(
				http.StatusOK,
				map[string]string{"status":"warning","contents":"The module information may not have changed"},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{"status":"refresh","contents":"module modify successfully~"},
		)

	}

}

