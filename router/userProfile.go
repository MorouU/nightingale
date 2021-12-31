package appRouter

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	appConfig "nightingale/config"
	lib "nightingale/lib/func"
	sqli "nightingale/lib/sql"
	moduleLib "nightingale/module"
	publicModule "nightingale/module/public"
	"html"
	"html/template"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// 用户个人
func pages_userProfile_profile(c *gin.Context){

	// CHECK LOGIN
	if !func_userProfile_isLogin(c) || !func_userProfile_checkAuthID(c){
		sessionClear(c)
		c.Redirect(
				http.StatusFound,
				"/userEnter/login",
			)
		return
	}
	// SHOW
	if c.Request.Method == "GET"{
		authID := sessionGet(c,"authID").(string)
		userName := sessionGet(c,"userName").(string)
		userEmail := sessionGet(c,"userEmail").(string)
		// 连接数据库
		thisSql := &sqli.Sql{}
		thisSql.Connect()
		defer thisSql.Close()
		// 获取用户数据
		r,_ := thisSql.Query("userLogin", []string{"userPhone","userLevel"}, []string{"userName","userEmail"},
				[]string{userName,userEmail}, " AND ", "",
			)
		// 若获取不到则返回认证出错
		if len(r) != 1{
			sessionClear(c)
			c.Redirect(
				http.StatusFound,
				"/userEnter/login",
			)
		}
		// 若获取到则将渲染内容回显至模板页面
		userPhone := r[0]["userPhone"].(string)
		userLevel := appConfig.UserLevelString[r[0]["userLevel"].(int64)]
		c.HTML(http.StatusOK, "userProfile_useredit.ni",
			gin.H{"authID": authID, "userName":userName, "userEmail":userEmail, "userPhone":userPhone, "userLevel":userLevel,
			}, )
	}

	// CHANGE
	if c.Request.Method == "POST"{
		func_userProfile_changeInfo(c)
	}

}

// 用户模块
func pages_userProfile_modules(c *gin.Context){

	// CHECK LOGIN
	if !func_userProfile_isLogin(c) || !func_userProfile_checkAuthID(c){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/userEnter/login",
		)
		return
	}

	// SHOW
	if c.Request.Method == "GET"{

		authID := sessionGet(c,"authID").(string)
		userID := strconv.FormatInt(sessionGet(c,"userID").(int64),10)
		customModules := ""
		appModules := ""
		/*
			SQL
		*/
		// Connect
		thisSql := &sqli.Sql{}
		thisSql.Connect()
		defer thisSql.Close()
		// Get user custom modules
		r,_:=thisSql.Query(
				"userModule",
				[]string{"*"},
				[]string{"belongUserID"},
				[]string{userID},
				" AND ",
				"",
			)
		if len(r) > 0 {
			for _,block := range r {
				customModules += fmt.Sprintf(moduleLib.UserCustomModulesFormat,
					block["moduleID"].(int64),
					block["moduleID"].(int64),
					block["moduleID"].(int64),
					html.EscapeString(block["moduleName"].(string)),
					block["moduleID"].(int64),
					html.EscapeString(block["moduleData"].(string)),
					block["moduleID"].(int64),
					html.EscapeString(block["createTime"].(string)),
					)

			}
		}

		// Get app public modules

		rt,_ := thisSql.Query(
				"appModule",
				[]string{"*"},
				[]string{},
				[]string{},
				"",
				"",
			)
		if len(rt) > 0 {
			for _,block := range rt {

				appModules += fmt.Sprintf(moduleLib.AppPublicModulesFormat,
					block["moduleID"].(string),
					block["moduleID"].(string),
					block["moduleID"].(string),
					html.EscapeString(block["moduleName"].(string)),
					block["moduleID"].(string),
					html.EscapeString(block["moduleData"].(string)),
					block["moduleID"].(string),
					html.EscapeString(block["createTime"].(string)),
				)

			}
		}

		c.HTML(
			http.StatusOK,
			"userProfile_modules.ni",
			gin.H{
				"authID":authID,
				"appModules":template.HTML(appModules),
				"userCustomModules":template.HTML(customModules),
			},
			)
	}

	// CHANGE

	if c.Request.Method == "POST"{
		func_userProfile_controlModules(c)
	}

}

// 用户主页
func pages_userProfile_home(c *gin.Context){

	// CHECK LOGIN
	if !func_userProfile_isLogin(c) || !func_userProfile_checkAuthID(c){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/userEnter/login",
		)
		return
	}
	// SHOW
	if c.Request.Method == "GET"{
		authID := sessionGet(c,"authID").(string)
		userID := strconv.FormatInt(sessionGet(c,"userID").(int64),10)
		taskRow := "";pagesTotal := 1;enableTask := 0;disableTask := 0;taskStatus := ""
		// 连接数据库
		thisSql := &sqli.Sql{}
		thisSql.Connect()
		defer thisSql.Close()
		// 获取任务总览
		r,_ := thisSql.Query("appTasks", []string{"*"}, []string{"taskUserID"}, []string{userID}, " AND ",
		"", )
		total := len(r)
		// 规划每页显示内容
		showTotal := appConfig.TaskShowPerPage
		if total > 0 {
			// 统计任务启用/停用个数
			for _,row := range r{
				if row["taskStatus"].(int64) == 0{disableTask += 1} else{enableTask += 1}
			}
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.TaskShowPerPage)))
			if total < appConfig.TaskShowPerPage {showTotal = total}
			for _,row := range r[:showTotal]{
				if row["taskStatus"].(int64) == 0{taskStatus = "停用"} else{taskStatus = "启用"}
				// 渲染模板页面
				taskRow += fmt.Sprintf(moduleLib.TaskTableHomeFormat, row["taskID"].(int64),
					html.EscapeString(row["taskName"].(string)), html.EscapeString(row["taskModuleName"].(string)),
					row["taskRecordNum"].(int64), taskStatus, html.EscapeString(row["taskParams"].(string)),
					html.EscapeString(row["taskCreateTime"].(string)), )
			}}
		// 将渲染内容回显至模板页面
		c.HTML(http.StatusOK, "userProfile_home.ni",
			gin.H{"authID":authID, "taskTotal":total, "enableTaskNum":enableTask, "disableTaskNum":disableTask,
				"taskRow":template.HTML(taskRow), "pagesTotal":pagesTotal,}, )
	}

	// FIND
	if c.Request.Method == "POST"{

		func_userProfile_controlHome(c)

	}
}

// 用户任务
func pages_userProfile_tasks(c *gin.Context){

	// CHECK LOGIN
	if !func_userProfile_isLogin(c) || !func_userProfile_checkAuthID(c){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/userEnter/login",
		)
		return
	}

	// SHOW
	if c.Request.Method == "GET"{
		authID := sessionGet(c,"authID").(string)
		userID := strconv.FormatInt(sessionGet(c,"userID").(int64),10)
		taskRow := "";pagesTotal := 1;taskStatus := "";taskControl := "";moduleUserOption := "";modulePublicOption := ""
		// 连接数据库
		thisSql := &sqli.Sql{}
		thisSql.Connect()
		defer thisSql.Close()
		// 获取任务内容
		r,_ := thisSql.Query(
			"appTasks",
			[]string{"*"}, []string{"taskUserID"}, []string{userID}, " AND ", "", )
		// 处理任务状态内容
		total := len(r)
		showTotal := appConfig.TaskShowPerPage
		if total > 0 {
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.TaskShowPerPage)))
			if total < appConfig.TaskShowPerPage {showTotal = total}
			for _,row := range r[:showTotal]{
				if row["taskStatus"].(int64) == 0{
					taskStatus = "停用"
				taskControl = "启用"
				} else{
					taskStatus = "启用"
				taskControl = "停用"}
				// 渲染模板页面
				taskRow += fmt.Sprintf(moduleLib.TaskTableTasksFormat,
					row["taskID"].(int64), html.EscapeString(row["taskName"].(string)),
					html.EscapeString(row["taskModuleName"].(string)), row["taskRecordNum"].(int64),
					row["taskID"].(int64), taskStatus, html.EscapeString(row["taskParams"].(string)),
					html.EscapeString(row["taskCreateTime"].(string)), row["taskID"].(int64),
					row["taskID"].(int64), html.EscapeString(taskControl), row["taskID"].(int64), row["taskID"].(int64),
					row["taskID"].(int64), row["taskID"].(int64), row["taskID"].(int64), row["taskID"].(int64),
				)
			}
		}
		// get modules
		rt,_ := thisSql.Query(
				"userModule",
				[]string{"moduleID","moduleName"}, []string{"belongUserID"}, []string{userID}, " AND ", "",
			)


		if len(rt) > 0{
			for _, row := range rt {
				moduleUserOption += fmt.Sprintf(moduleLib.TaskUserOptionTasksFormat, row["moduleID"].(int64),
						row["moduleID"].(int64), html.EscapeString(row["moduleName"].(string)),
					)
			}
		}
		rts,_ := thisSql.Query(
			"appModule",
			[]string{"moduleID","moduleName"},
			[]string{},
			[]string{},
			" AND ",
			"",
		)
		if len(rts) > 0{
			for _, row := range rts {

				modulePublicOption += fmt.Sprintf(moduleLib.TaskPublicOptionTasksFormat,
					row["moduleID"].(string),
					row["moduleID"].(string),
					html.EscapeString(row["moduleName"].(string)),
				)

			}
		}
		c.HTML(
			http.StatusOK,
			"userProfile_tasks.ni",
			gin.H{
				"authID":authID,
				"personModule":template.HTML(moduleUserOption),
				"publicModule":template.HTML(modulePublicOption),
				"taskRow":template.HTML(taskRow),
				"pagesTotal":pagesTotal,
			},
		)

	}

	// ACTION

	if c.Request.Method == "POST"{

		func_userProfile_controlTasks(c)

	}

}

// 用户个人一览
func pages_userProfile_view(c *gin.Context){

	// CHECK LOGIN
	if !func_userProfile_isLogin(c) || !func_userProfile_checkAuthID(c){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/userEnter/login",
		)
		return
	}

	// SHOW
	if c.Request.Method == "GET"{

		authID := sessionGet(c,"authID").(string)
		userID := strconv.FormatInt(sessionGet(c,"userID").(int64),10)
		taskID := c.Param("taskID")
		taskRow := ""
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
			"appTaskView",
			[]string{"*"},
			[]string{"taskID","userID"},
			[]string{taskID,userID},
			" AND ",
			"",
		)

		total := len(r)
		showTotal := appConfig.RecordShowPerPage
		if total > 0 {

			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.RecordShowPerPage)))
			if total < appConfig.RecordShowPerPage {
				showTotal = total
			}

			for _,row := range r[:showTotal]{


				taskRow += fmt.Sprintf(moduleLib.TaskViewTasksFormat,
					row["recordID"].(int64),
					html.EscapeString(row["getMethod"].(string)),
					html.EscapeString(row["getIP"].(string)),
					html.EscapeString(row["getTime"].(string)),
					row["recordID"].(int64),
					row["recordID"].(int64),
					row["recordID"].(int64),
					row["recordID"].(int64),

				)

			}

		}

		// get task data

		rt,_ := thisSql.Query(
				"appTasks",
				[]string{"*"},
				[]string{"taskUserID","taskID"},
				[]string{userID,taskID},
				" AND ",
				"",
			)

		if len(rt) != 1{

			c.Redirect(
				http.StatusFound,
				"/userEnter/"+userID+"/tasks",
			)

		}

		c.HTML(
			http.StatusOK,
			"userProfile_record.ni",
			gin.H{
				"authID":authID,
				"taskID":taskID,
				"taskName":html.EscapeString(rt[0]["taskName"].(string)),
				"taskData":html.EscapeString(rt[0]["taskData"].(string)),
				"taskRow":template.HTML(taskRow),
				"pagesTotal":pagesTotal,
			},
		)

	}

	// VIEW
	if c.Request.Method == "POST"{

		func_userProfile_controlView(c)

	}


}

// 用户登出
func pages_userProfile_logout(c *gin.Context){

	// CHECK LOGIN
	if !func_userProfile_isLogin(c) || !func_userProfile_checkAuthID(c){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/userEnter/login",
		)
		return
	}

	sessionClear(c)
	c.Redirect(
		http.StatusFound,
		"/userEnter/login",
	)

}

/*
	userProfile
 */
// 从session判断是否登录
func func_userProfile_isLogin(c *gin.Context) bool {

	return sessionKeyExist(sessionGet(c,"isLogin"),"string") && sessionGet(c,"isLogin").(string) == "1"

}
// 从session判断认证ID是否合法
func func_userProfile_checkAuthID(c *gin.Context) bool{

	return sessionKeyExist(sessionGet(c,"authID"),"string") && c.Param("uuid") == sessionGet(c,"authID").(string)

}

// 从请求判断请求内容是与认证ID一致
func func_userProfile_checkAuthIDFromRequest(c *gin.Context,method,key string) bool{

	switch method {
	case	"GET":
		return sessionKeyExist(sessionGet(c,"authID"),"string") && c.DefaultQuery(key,"NoKey") == sessionGet(c,"authID").(string)
	default:
		return sessionKeyExist(sessionGet(c,"authID"),"string") && c.DefaultPostForm(key,"NoKey") == sessionGet(c,"authID").(string)
	}
}

// 用户信息更改
func func_userProfile_changeInfo(c *gin.Context){

	// 检查用户登录状态
	if !func_userProfile_isLogin(c) || !func_userProfile_checkAuthID(c) || !func_userProfile_checkAuthIDFromRequest(c,
		"POST","authID"){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/userEnter/login",
		)
		return
	}

	// 检查传入参数
	step := c.DefaultPostForm("step","")
	// 若不合法则返回错误
	if step != "info" && step != "pass"{
		c.JSON(http.StatusOK, map[string]string{"status":"error", "contents":"step error",}, )
		return
	}
	userName := sessionGet(c,"userName").(string)
	userEmail := sessionGet(c,"userEmail").(string)
	// 连接数据库
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()
	// 对执行步骤进行判断
	switch step {
	// 若为执行步骤为info则进行用户信息修改操作
	case "info":
		editUserName := c.DefaultPostForm("editUserName","")
		editUserEmail := c.DefaultPostForm("editUserEmail","")
		editUserPhone := c.DefaultPostForm("editUserPhone","")


		// Check Username
		if !lib.CheckUsername(editUserName){
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
		if !lib.VerifyEmailFormat(editUserEmail){
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
		if len(editUserPhone) > 0 &&!lib.VerifyMobileFormat(editUserPhone){
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
			"userLogin",
			[]string{"userName","userEmail","userPhone"},
			[]string{editUserName,editUserEmail,editUserPhone},
			[]string{"userName","userEmail"},
			[]string{userName,userEmail},
			" AND ",
			"",
			)

		if r == 0{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"warning",
					"contents":"The information may not have changed",
					"userName":editUserName,
					"userEmail":editUserEmail,
					"userPhone":editUserPhone,
				},
			)
			return
		}
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"contents":"Successful change~",
				"userName":editUserName,
				"userEmail":editUserEmail,
				"userPhone":editUserPhone,
			},
		)
	// 若执行步骤为pass则进行用户密码修改操作
	case "pass":

		changeUserCPass := c.DefaultPostForm("changeUserCPass","")
		changeUserNPass := c.DefaultPostForm("changeUserNPass","")
		changeUserVPass := c.DefaultPostForm("changeUserVPass","")

		// Check New pass
		if changeUserNPass != changeUserVPass{
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
				"userLogin",
				[]string{"userPasswd","userSalt"},
				[]string{"userName","userEmail"},
				[]string{userName,userEmail},
				" AND ",
				"",
			)
		// ?
		if len(r) != 1{
			sessionClear(c)
			c.Redirect(
				http.StatusFound,
				"/userEnter/login",
			)
		}
		cPassHash := lib.GetHashSignature([]byte(changeUserCPass),[]byte(r[0]["userSalt"].(string)),"sha256",true).(string)
		// Cmp Pass hash
		if cPassHash != r[0]["userPasswd"].(string){
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
		nPassHash := lib.GetHashSignature([]byte(changeUserNPass),[]byte(nSalt),"sha256",true).(string)

		rt,_ := thisSql.Update(
				"userLogin",
				[]string{"userPasswd","userSalt"},
				[]string{nPassHash,nSalt},
				[]string{"userName","userEmail"},
				[]string{userName,userEmail},
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
	userModules
 */
// 用户模块操作
func func_userProfile_controlModules(c *gin.Context){

	// CHECK LOGIN
	if !func_userProfile_isLogin(c) || !func_userProfile_checkAuthID(c) || !func_userProfile_checkAuthIDFromRequest(c,"POST","authID"){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/userEnter/login",
		)
		return
	}

	// check param
	step := c.DefaultPostForm("step","")

	if step != "add" && step != "modify" && step != "delete" && step != "view"{
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"error",
				"contents":"step error",
			},
		)
		return
	}

	userID := strconv.FormatInt(sessionGet(c,"userID").(int64),10)


	/*
		SQL
	*/
	// Connect
	thisSql := &sqli.Sql{}
	thisSql.Connect()
	defer thisSql.Close()

	switch step {
	case "add":
		moduleName := c.DefaultPostForm("moduleAddName","")
		moduleData := c.DefaultPostForm("moduleAddData","")
		moduleCode := c.DefaultPostForm("moduleAddCode","")

		// Check user Level

		rt,_ := thisSql.Query(
				"userLogin",
				[]string{"userLevel"},
				[]string{"userID"},
				[]string{userID},
				" AND ",
				"",
			)

		if len(rt) != 1 || rt[0]["userLevel"].(int64) < 2{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"This user could not create a custom module",
				},
			)
			return
		}

		// Check module name
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

		// insert
		r,_ := thisSql.Insert(
				"userModule",
				[]string{"moduleID","moduleName","moduleData","moduleCode","belongUserID","createTime"},
				[]string{"0",moduleName,moduleData,moduleCode,userID, lib.GetTimeDateTime()},
				"",
			)
		if r <=0 {
			if ! lib.CheckCustomModuleName(moduleName){
				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"error",
						"contents":"Added module failed",
					},
				)
				return
			}
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"Added module success~",
			},
		)

	case "modify":

		moduleID := c.DefaultPostForm("moduleID","")
		moduleName := c.DefaultPostForm("moduleModifyName","")
		moduleData := c.DefaultPostForm("moduleModifyData","")
		moduleCode := c.DefaultPostForm("moduleModifyCode","")

		// Check module name
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
				map[string]string{
					"status":"warning",
					"contents":"The module has not been modified",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"Modify module successfully~",
			},
		)

	case "delete":

		moduleID := c.DefaultPostForm("moduleID","")

		// delete
		r,_ :=thisSql.Delete(
				"userModule",
				[]string{"moduleID","belongUserID"},
				[]string{moduleID,userID},
				" AND ",
				"",
			)

		if r <= 0{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Failed to delete module",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"Delete module successfully~",
			},
		)

	case "view":

		moduleID := c.DefaultPostForm("moduleID","")

		// view
		r,_ := thisSql.Query(
				"userModule",
				[]string{"moduleCode"},
				[]string{"moduleID","belongUserID"},
				[]string{moduleID,userID},
				" AND ",
				"",
			)

		if len(r) > 0{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"ok",
					"code":r[0]["moduleCode"].(string),
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"error",
				"contents":"Failed to read code",
			},
		)
		return

	}

}

/*
	userHome
 */
// 用户主页操作
func func_userProfile_controlHome(c *gin.Context){

	// CHECK LOGIN
	if !func_userProfile_isLogin(c) || !func_userProfile_checkAuthID(c) || !func_userProfile_checkAuthIDFromRequest(c,"POST","authID"){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/userEnter/login",
		)
		return
	}

	// check param
	step := c.DefaultPostForm("step","")

	if step != "previous" && step != "next" && step != "back" && step != "find"{
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"error",
				"contents":"step error",
			},
		)
		return
	}

	userID := strconv.FormatInt(sessionGet(c,"userID").(int64),10)

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

			endPages := (currentPage - 1) * appConfig.TaskShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.TaskShowPerPage)))

			if currentPage - 1 > 0 && currentPage - 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage - 2)*appConfig.TaskShowPerPage :endPages]{


					if row["taskStatus"].(int64) == 0{
						taskStatus = "停用"
					} else{
						taskStatus = "启用"
					}


					taskRow += fmt.Sprintf(moduleLib.TaskTableHomeFormat,
						row["taskID"].(int64),
						html.EscapeString(row["taskName"].(string)),
						html.EscapeString(row["taskModuleName"].(string)),
						row["taskRecordNum"].(int64),
						taskStatus,
						html.EscapeString(row["taskParams"].(string)),
						html.EscapeString(row["taskCreateTime"].(string)),
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

			endPages := (currentPage + 1) * appConfig.TaskShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.TaskShowPerPage)))

			if currentPage + 1 > 1 && currentPage + 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage)*appConfig.TaskShowPerPage :endPages]{


					if row["taskStatus"].(int64) == 0{
						taskStatus = "停用"
					} else{
						taskStatus = "启用"
					}


					taskRow += fmt.Sprintf(moduleLib.TaskTableHomeFormat,
						row["taskID"].(int64),
						html.EscapeString(row["taskName"].(string)),
						html.EscapeString(row["taskModuleName"].(string)),
						row["taskRecordNum"].(int64),
						taskStatus,
						html.EscapeString(row["taskParams"].(string)),
						html.EscapeString(row["taskCreateTime"].(string)),
					)

				}

				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"html":taskRow,
						"page":strconv.Itoa(currentPage + 1),
						"total": strconv.Itoa(pagesTotal),
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
		showTotal := appConfig.TaskShowPerPage
		if total > 0 {

			pagesTotal = int(math.Ceil(float64(total) / float64(appConfig.TaskShowPerPage)))
			if total < appConfig.TaskShowPerPage {
				showTotal = total
			}

			for _, row := range r[:showTotal] {

				if row["taskStatus"].(int64) == 0 {
					taskStatus = "停用"
				} else {
					taskStatus = "启用"
				}

				taskRow += fmt.Sprintf(moduleLib.TaskTableHomeFormat,
					row["taskID"].(int64),
					html.EscapeString(row["taskName"].(string)),
					html.EscapeString(row["taskModuleName"].(string)),
					row["taskRecordNum"].(int64),
					taskStatus,
					html.EscapeString(row["taskParams"].(string)),
					html.EscapeString(row["taskCreateTime"].(string)),
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

		taskID := c.DefaultPostForm("findTaskID","")
		taskRow := ""
		taskStatus := ""


		// get total
		r,_ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskID","taskUserID"},
			[]string{taskID,userID},
			" AND ",
			"",
		)


		if len(r) == 1{

			if r[0]["taskStatus"].(int64) == 0 {
				taskStatus = "停用"
			} else {
				taskStatus = "启用"
			}

			taskRow = fmt.Sprintf(moduleLib.TaskTableHomeFormat,
				r[0]["taskID"].(int64),
				html.EscapeString(r[0]["taskName"].(string)),
				html.EscapeString(r[0]["taskModuleName"].(string)),
				r[0]["taskRecordNum"].(int64),
				taskStatus,
				html.EscapeString(r[0]["taskParams"].(string)),
				html.EscapeString(r[0]["taskCreateTime"].(string)),
			)

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

	}

}

/*
	userTasks
 */
// 用户任务操作
func func_userProfile_controlTasks(c *gin.Context){

	// CHECK LOGIN
	if !func_userProfile_isLogin(c) || !func_userProfile_checkAuthID(c) || !func_userProfile_checkAuthIDFromRequest(c,"POST","authID"){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/userEnter/login",
		)
		return
	}

	// check param
	step := c.DefaultPostForm("step","")

	if step != "previous" && step != "next" && step != "back" && step != "find" && step != "control" && step != "delete" && step != "getApi" && step != "new" && step != "customAdd" && step != "defaultAdd" && step != "requestAdd" && step != "fishAdd" && step != "portAdd" && step!="viewTask"{
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"error",
				"contents":"step error",
			},
		)
		return
	}

	userID := strconv.FormatInt(sessionGet(c,"userID").(int64),10)

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
		taskControl := ""
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

			endPages := (currentPage - 1) * appConfig.TaskShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.TaskShowPerPage)))

			if currentPage - 1 > 0 && currentPage - 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage - 2)*appConfig.TaskShowPerPage :endPages]{


					if row["taskStatus"].(int64) == 0{
						taskStatus = "停用"
						taskControl = "启用"
					} else{
						taskStatus = "启用"
						taskControl = "停用"
					}


					taskRow += fmt.Sprintf(moduleLib.TaskTableTasksFormat,
						row["taskID"].(int64),
						html.EscapeString(row["taskName"].(string)),
						html.EscapeString(row["taskModuleName"].(string)),
						row["taskRecordNum"].(int64),
						row["taskID"].(int64),
						taskStatus,
						html.EscapeString(row["taskParams"].(string)),
						html.EscapeString(row["taskCreateTime"].(string)),

						row["taskID"].(int64),
						row["taskID"].(int64),
						html.EscapeString(taskControl),
						row["taskID"].(int64),
						row["taskID"].(int64),
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
		taskControl := ""
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

			endPages := (currentPage + 1) * appConfig.TaskShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.TaskShowPerPage)))

			if currentPage + 1 > 1 && currentPage + 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage)*appConfig.TaskShowPerPage :endPages]{


					if row["taskStatus"].(int64) == 0{
						taskStatus = "停用"
						taskControl = "启用"
					} else{
						taskStatus = "启用"
						taskControl = "停用"
					}


					taskRow += fmt.Sprintf(moduleLib.TaskTableTasksFormat,
						row["taskID"].(int64),
						html.EscapeString(row["taskName"].(string)),
						html.EscapeString(row["taskModuleName"].(string)),
						row["taskRecordNum"].(int64),
						row["taskID"].(int64),
						taskStatus,
						html.EscapeString(row["taskParams"].(string)),
						html.EscapeString(row["taskCreateTime"].(string)),

						row["taskID"].(int64),
						row["taskID"].(int64),
						html.EscapeString(taskControl),
						row["taskID"].(int64),
						row["taskID"].(int64),
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
		taskControl := ""

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
		showTotal := appConfig.TaskShowPerPage
		if total > 0 {

			pagesTotal = int(math.Ceil(float64(total) / float64(appConfig.TaskShowPerPage)))
			if total < appConfig.TaskShowPerPage {
				showTotal = total
			}

			for _, row := range r[:showTotal] {

				if row["taskStatus"].(int64) == 0{
					taskStatus = "停用"
					taskControl = "启用"
				} else{
					taskStatus = "启用"
					taskControl = "停用"
				}

				taskRow += fmt.Sprintf(moduleLib.TaskTableTasksFormat,
					row["taskID"].(int64),
					html.EscapeString(row["taskName"].(string)),
					html.EscapeString(row["taskModuleName"].(string)),
					row["taskRecordNum"].(int64),
					row["taskID"].(int64),
					taskStatus,
					html.EscapeString(row["taskParams"].(string)),
					html.EscapeString(row["taskCreateTime"].(string)),

					row["taskID"].(int64),
					row["taskID"].(int64),
					html.EscapeString(taskControl),
					row["taskID"].(int64),
					row["taskID"].(int64),
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

		taskID := c.DefaultPostForm("findTaskID","")
		taskRow := ""
		taskStatus := ""
		taskControl := ""


		// get total
		r,_ := thisSql.Query(
			"appTasks",
			[]string{"*"},
			[]string{"taskID","taskUserID"},
			[]string{taskID,userID},
			" AND ",
			"",
		)


		if len(r) == 1{

			if r[0]["taskStatus"].(int64) == 0{
				taskStatus = "停用"
				taskControl = "启用"
			} else{
				taskStatus = "启用"
				taskControl = "停用"
			}

			taskRow += fmt.Sprintf(moduleLib.TaskTableTasksFormat,
				r[0]["taskID"].(int64),
				html.EscapeString(r[0]["taskName"].(string)),
				html.EscapeString(r[0]["taskModuleName"].(string)),
				r[0]["taskRecordNum"].(int64),
				r[0]["taskID"].(int64),
				taskStatus,
				html.EscapeString(r[0]["taskParams"].(string)),
				html.EscapeString(r[0]["taskCreateTime"].(string)),

				r[0]["taskID"].(int64),
				r[0]["taskID"].(int64),
				html.EscapeString(taskControl),
				r[0]["taskID"].(int64),
				r[0]["taskID"].(int64),
				r[0]["taskID"].(int64),
				r[0]["taskID"].(int64),
				r[0]["taskID"].(int64),
				r[0]["taskID"].(int64),
			)

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

	case "control":

		taskID := c.DefaultPostForm("taskID","")
		taskStatus := "0"

		// Get status
		r,_ := thisSql.Query(
				"appTasks",
				[]string{"taskStatus"},
				[]string{"taskID","taskUserID"},
				[]string{taskID,userID},
				" AND ",
				"",
			)

		if len(r) != 1{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Task ID error",
				},
			)
			return
		}
		if r[0]["taskStatus"].(int64) != 1{
			taskStatus = "1"
		}
		// update
		rt,_ := thisSql.Update(
				"appTasks",
				[]string{"taskStatus"},
				[]string{taskStatus},
				[]string{"taskID","taskUserID"},
				[]string{taskID,userID},
				" AND ",
				"",
			)

		if rt != 1{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"warning",
					"contents":"The task state may not have changed",
				},
			)
			return

		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"Change task status successfully~",
			},
		)


	case "delete":

		taskID := c.DefaultPostForm("taskID","")

		// delete

		rt,_ := thisSql.Query(
			"appTasks",
			[]string{"filePath"},
			[]string{"taskUserID","taskID"},
			[]string{userID,taskID},
			" AND ",
			"",
		)
		if len(rt) > 0{
			for _,row := range strings.Split(rt[0]["filePath"].(string),"|"){
				os.Remove(row)
			}
		}

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
				[]string{"userID","taskID"},
				[]string{userID,taskID},
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
	case "getApi":
		// 获取任务ID
		taskID := c.DefaultPostForm("taskID","")
		// 从任务ID获取API
		r,_ := thisSql.Query(
				"appTasks",
				[]string{"taskName","taskApi"},
				[]string{"taskID","taskUserID"},
				[]string{taskID,userID},
				" AND ",
				"",
			)
		// 若无法获取到API则返回错误信息
		if len(r) != 1{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Can not get task api",
				},
			)
			return
		}
		// 通过JSON返回有关内容
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"taskName":r[0]["taskName"].(string),
				"taskApi":fmt.Sprintf(appConfig.ApiDomain, r[0]["taskApi"].(string)),
			},
		)

	case "new":

		taskName := c.DefaultPostForm("taskName","")
		taskModule := c.DefaultPostForm("taskModule","")
		tableName := ""
		taskModuleID := ""


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

		tableName = lib.VerifyModuleType(taskModule)
		if tableName  == "error"{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"The task module is not valid",
				},
			)
			return
		}

		taskModuleID = strings.Split(taskModule,"_")[1]

		// Get data
		if tableName == "user"{

			// Get code
			r,_ := thisSql.Query(
					"userModule",
					[]string{"moduleName","moduleCode"},
					[]string{"moduleID","belongUserID",},
					[]string{taskModuleID,userID},
					" AND ",
					"",
				)

			if len(r) != 1{
				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"error",
						"contents":"Module error",
					},
				)
				return
			}

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"nextCustom",
					"taskName":r[0]["moduleName"].(string),
					"taskCode":r[0]["moduleCode"].(string),
					"moduleID":taskModuleID,
				},
			)

		}

		if tableName == "public" {

			// Get module
			r,_ := thisSql.Query(
					"appModule",
					[]string{"moduleName"},
					[]string{"moduleID"},
					[]string{taskModuleID},
					" AND ",
					"",
				)
			if len(r) != 1{
				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"error",
						"contents":"Module error",
					},
				)
				return
			}

			switch r[0]["moduleName"] {
			case publicModule.AllPublicModuleName["default"]:

				publicModule.Module_default_next(c,taskModuleID,taskName)

			case publicModule.AllPublicModuleName["request"]:

				publicModule.Module_request_next(c,taskModuleID,taskName)

			case publicModule.AllPublicModuleName["fish"]:

				publicModule.Module_fish_next(c,taskModuleID,taskName)

			case publicModule.AllPublicModuleName["port"]:

				publicModule.Module_port_next(c,taskModuleID,taskName)

			}


		}

	case "customAdd":

		taskName := c.DefaultPostForm("taskName","")
		taskData := c.DefaultPostForm("taskData","")
		taskStatus := c.DefaultPostForm("taskStatus","")
		taskCode := c.DefaultPostForm("taskCode","")
		moduleID := c.DefaultPostForm("moduleID","")

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

		if taskStatus != "1" && taskStatus != "0"{
			taskStatus = "0"
		}

		// Get data

		r,_ := thisSql.Query(
				"userModule",
				[]string{"moduleName"},
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
					"contents":"Module error",
				},
			)
			return
		}

		// insert

		rt,_ := thisSql.Insert(
				"appTasks",
				[]string{"taskID",
					"taskModuleID",
					"taskUserID",
					"taskName",
					"taskModuleName",
					"taskModulePublic",
					"taskData",
					"taskCode",
					"taskApi",
					"taskParams",
					"taskStatus",
					"taskRecordNum",
					"filePath",
					"taskCreateTime"},
				[]string{"0",
					moduleID,
					userID,
					taskName,
					r[0]["moduleName"].(string),
					"0",
					taskData,
					taskCode,
					lib.GetHashSignature([]byte(taskName),[]byte(uuid.New()),"sha256",true).(string),
					"",
					taskStatus,
					"0",
					"",
					lib.GetTimeDateTime(),
				},
				"",
			)

		if rt != 1{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"New task failed",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"New task successfully created~",
			},
		)

	case "defaultAdd":

		taskName := c.DefaultPostForm("taskName","")
		taskData := c.DefaultPostForm("taskData","")
		taskStatus := c.DefaultPostForm("taskStatus","")
		moduleID := c.DefaultPostForm("moduleID","")

		currentUrl := c.DefaultPostForm("defaultCurrentUrl","false")
		refererUrl := c.DefaultPostForm("defaultRefererUrl","false")
		cookie := c.DefaultPostForm("defaultCookie","false")
		OS := c.DefaultPostForm("defaultOS","false")
		browser := c.DefaultPostForm("defaultBrowser","false")
		screenResolution := c.DefaultPostForm("defaultScreenResolution","false")
		webPage := c.DefaultPostForm("defaultWebPage","false")
		screenShot := c.DefaultPostForm("defaultScreenShot","false")

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

		if taskStatus != "1" && taskStatus != "0"{
			taskStatus = "0"
		}

		jsonData,_ := json.Marshal(&publicModule.DefaultStruct{
			func_userProfile_defaultModuleParamsCheck(currentUrl),
			func_userProfile_defaultModuleParamsCheck(refererUrl),
			func_userProfile_defaultModuleParamsCheck(cookie),
			func_userProfile_defaultModuleParamsCheck(OS),
			func_userProfile_defaultModuleParamsCheck(browser),
			func_userProfile_defaultModuleParamsCheck(screenResolution),
			func_userProfile_defaultModuleParamsCheck(webPage),
			func_userProfile_defaultModuleParamsCheck(screenShot),
		},
		)

		// Get data

		r,_ := thisSql.Query(
			"appModule",
			[]string{"moduleName"},
			[]string{"moduleID"},
			[]string{moduleID},
			" AND ",
			"",
		)

		if len(r) != 1{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Module error",
				},
			)
			return
		}

		// insert

		rt,_ := thisSql.Insert(
			"appTasks",
			[]string{"taskID",
				"taskModuleID",
				"taskUserID",
				"taskName",
				"taskModuleName",
				"taskModulePublic",
				"taskData",
				"taskCode",
				"taskApi",
				"taskParams",
				"taskStatus",
				"taskRecordNum",
				"filePath",
				"taskCreateTime"},
			[]string{"0",
				moduleID,
				userID,
				taskName,
				r[0]["moduleName"].(string),
				"1",
				taskData,
				"",
				lib.GetHashSignature([]byte(taskName),[]byte(uuid.New()),"sha256",true).(string),
				string(jsonData),
				taskStatus,
				"0",
				"",
				lib.GetTimeDateTime(),
			},
			"",
		)

		if rt != 1{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"New task failed",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"New task successfully created~",
			},
		)

	case "requestAdd":

		taskName := c.DefaultPostForm("taskName","")
		taskData := c.DefaultPostForm("taskData","")
		taskStatus := c.DefaultPostForm("taskStatus","")
		moduleID := c.DefaultPostForm("moduleID","")

		rUrl := c.DefaultPostForm("requestUrl","")
		rGET := c.DefaultPostForm("requestGET","")
		rPOST := c.DefaultPostForm("requestPOST","")
		form,_ := c.MultipartForm()
		rFile := form.File["requestFILE"]
		filePrefix := lib.GetHashSignature([]byte(userID + "_"),[]byte(userID),"md5",true).(string)
		fileNames := make([]string,0,len(rFile))
		paths := make([]string,0,len(rFile))

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

		if taskStatus != "1" && taskStatus != "0"{
			taskStatus = "0"
		}


		// Get data

		r,_ := thisSql.Query(
			"appModule",
			[]string{"moduleName"},
			[]string{"moduleID"},
			[]string{moduleID},
			" AND ",
			"",
		)

		if len(r) != 1{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Module error",
				},
			)
			return
		}

		// insert

		for _,each := range rFile{
			fileName := filePrefix +"_"+ lib.GetHashSignature([]byte(filepath.Base(each.Filename) + "_"),[]byte(lib.GetRandStr(10,"")),"sha256",true).(string)
			c.SaveUploadedFile(each, publicModule.RequestConfig["path"] +"_"+each.Filename+"_"+fileName)
			paths = append(paths, publicModule.RequestConfig["path"] +"_"+each.Filename+"_"+fileName)
			fileNames = append(fileNames,each.Filename)
		}

		jsonData,_ := json.Marshal(map[string]interface{}{
			"RequestUrl":rUrl,
			"RequestGET":rGET,
			"RequestPOST":rPOST,
			"RequestFILE":fileNames,
		})

		rt,_ := thisSql.Insert(
			"appTasks",
			[]string{"taskID",
				"taskModuleID",
				"taskUserID",
				"taskName",
				"taskModuleName",
				"taskModulePublic",
				"taskData",
				"taskCode",
				"taskApi",
				"taskParams",
				"taskStatus",
				"taskRecordNum",
				"filePath",
				"taskCreateTime"},
			[]string{"0",
				moduleID,
				userID,
				taskName,
				r[0]["moduleName"].(string),
				"1",
				taskData,
				"",
				lib.GetHashSignature([]byte(taskName),[]byte(uuid.New()),"sha256",true).(string),
				string(jsonData),
				taskStatus,
				"0",
				strings.Join(paths,"|"),
				lib.GetTimeDateTime(),
			},
			"",
		)

		if rt != 1{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"New task failed",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"New task successfully created~",
			},
		)

	case "portAdd":

		taskName := c.DefaultPostForm("taskName","")
		taskData := c.DefaultPostForm("taskData","")
		taskStatus := c.DefaultPostForm("taskStatus","")
		moduleID := c.DefaultPostForm("moduleID","")

		pI := c.DefaultPostForm("portIP","")
		pS := c.DefaultPostForm("portScan","")



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

		if taskStatus != "1" && taskStatus != "0"{
			taskStatus = "0"
		}


		// Get data

		r,_ := thisSql.Query(
			"appModule",
			[]string{"moduleName"},
			[]string{"moduleID"},
			[]string{moduleID},
			" AND ",
			"",
		)

		if len(r) != 1{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Module error",
				},
			)
			return
		}

		// insert


		jsonData,_ := json.Marshal(map[string]interface{}{
			"portIP":pI,
			"portScan":pS,
		})

		rt,_ := thisSql.Insert(
			"appTasks",
			[]string{"taskID",
				"taskModuleID",
				"taskUserID",
				"taskName",
				"taskModuleName",
				"taskModulePublic",
				"taskData",
				"taskCode",
				"taskApi",
				"taskParams",
				"taskStatus",
				"taskRecordNum",
				"filePath",
				"taskCreateTime"},
			[]string{"0",
				moduleID,
				userID,
				taskName,
				r[0]["moduleName"].(string),
				"1",
				taskData,
				"",
				lib.GetHashSignature([]byte(taskName),[]byte(uuid.New()),"sha256",true).(string),
				string(jsonData),
				taskStatus,
				"0",
				"",
				lib.GetTimeDateTime(),
			},
			"",
		)

		if rt != 1{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"New task failed",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"New task successfully created~",
			},
		)


	case "fishAdd":

		taskName := c.DefaultPostForm("taskName","")
		taskData := c.DefaultPostForm("taskData","")
		taskStatus := c.DefaultPostForm("taskStatus","")
		moduleID := c.DefaultPostForm("moduleID","")

		fUrl := c.DefaultPostForm("fishUrl","")
		fFile,err := c.FormFile("fishFILE")

		filePrefix := lib.GetHashSignature([]byte(userID + "_"),[]byte(userID),"md5",true).(string)
		fileName := ""
		filePath := ""


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

		if taskStatus != "1" && taskStatus != "0"{
			taskStatus = "0"
		}


		// Get data

		r,_ := thisSql.Query(
			"appModule",
			[]string{"moduleName"},
			[]string{"moduleID"},
			[]string{moduleID},
			" AND ",
			"",
		)

		if len(r) != 1{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Module error",
				},
			)
			return
		}

		// insert

		if err == nil {
			fileHashName := filePrefix + "_" + lib.GetHashSignature([]byte(filepath.Base(fFile.Filename)+"_"), []byte(lib.GetRandStr(10, "")), "sha256", true).(string)
			c.SaveUploadedFile(fFile, publicModule.FishConfig["path"]+"_"+fFile.Filename+"_"+fileHashName)
			filePath = publicModule.FishConfig["path"]+"_"+fFile.Filename+"_"+fileHashName
			fileName = fFile.Filename
		}


		jsonData,_ := json.Marshal(map[string]interface{}{
			"fishUrl":fUrl,
			"fishFILE":fileName,
		})

		rt,_ := thisSql.Insert(
			"appTasks",
			[]string{"taskID",
				"taskModuleID",
				"taskUserID",
				"taskName",
				"taskModuleName",
				"taskModulePublic",
				"taskData",
				"taskCode",
				"taskApi",
				"taskParams",
				"taskStatus",
				"taskRecordNum",
				"filePath",
				"taskCreateTime"},
			[]string{"0",
				moduleID,
				userID,
				taskName,
				r[0]["moduleName"].(string),
				"1",
				taskData,
				"",
				lib.GetHashSignature([]byte(taskName),[]byte(uuid.New()),"sha256",true).(string),
				string(jsonData),
				taskStatus,
				"0",
				filePath,
				lib.GetTimeDateTime(),
			},
			"",
		)

		if rt != 1{
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"New task failed",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"New task successfully created~",
			},
		)

	case "viewTask":

		taskID := c.DefaultPostForm("taskID","")

		// check
		r,_ := thisSql.Query(
				"appTasks",
				[]string{"*"},
				[]string{"taskID","taskUserID"},
				[]string{taskID,userID},
				" AND ",
				"",
			)

		if len(r) != 1{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Task ID error",
				},
			)
			return

		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"redirect",
				"url":taskID+"/view",
			},
		)

	}

}

// 强制转换
func func_userProfile_defaultModuleParamsCheck(s string) string{

	if s != "false" && s != "true"{
		return "false"
	}
	return s

}

/*
	userView
 */
// 用户个人一览操作
func func_userProfile_controlView(c *gin.Context){

	// CHECK LOGIN
	if !func_userProfile_isLogin(c) || !func_userProfile_checkAuthID(c) || !func_userProfile_checkAuthIDFromRequest(c,"POST","authID"){
		sessionClear(c)
		c.Redirect(
			http.StatusFound,
			"/userEnter/login",
		)
		return
	}

	// check param
	step := c.DefaultPostForm("step","")

	if step != "previous" && step != "next" && step != "back" && step != "find" && step != "delete" && step!="save" && step != "view" {
		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"error",
				"contents":"step error",
			},
		)
		return
	}

	userID := strconv.FormatInt(sessionGet(c,"userID").(int64),10)
	taskID := c.Param("taskID")

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
		viewRow := ""

		// get total
		r,_ := thisSql.Query(
			"appTaskView",
			[]string{"*"},
			[]string{"taskID","userID"},
			[]string{taskID,userID},
			" AND ",
			"",
		)

		total := len(r)
		if total > 0 {

			endPages := (currentPage - 1) * appConfig.RecordShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.RecordShowPerPage)))

			if currentPage - 1 > 0 && currentPage - 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage - 2)*appConfig.RecordShowPerPage :endPages]{


					viewRow += fmt.Sprintf(moduleLib.TaskViewTasksFormat,
						row["recordID"].(int64),
						html.EscapeString(row["getMethod"].(string)),
						html.EscapeString(row["getIP"].(string)),
						html.EscapeString(row["getTime"].(string)),
						row["recordID"].(int64),
						row["recordID"].(int64),
						row["recordID"].(int64),
						row["recordID"].(int64),

					)

				}

				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"html":viewRow,
						"page":strconv.Itoa(currentPage - 1),
						"total":strconv.Itoa(pagesTotal),
					},
				)

			}

		}



	case "next":

		currentPage,_ := strconv.Atoi(c.DefaultPostForm("currentPage",""))
		pagesTotal := 0
		viewRow := ""

		// get total
		r,_ := thisSql.Query(
			"appTaskView",
			[]string{"*"},
			[]string{"taskID","userID"},
			[]string{taskID,userID},
			" AND ",
			"",
		)

		total := len(r)
		if total > 0 {

			endPages := (currentPage + 1) * appConfig.RecordShowPerPage
			pagesTotal = int(math.Ceil(float64(total)/float64(appConfig.RecordShowPerPage)))

			if currentPage + 1 > 1 && currentPage + 1 <= pagesTotal{

				if endPages > total{
					endPages = total
				}

				for _,row := range r[(currentPage)*appConfig.RecordShowPerPage :endPages]{



					viewRow += fmt.Sprintf(moduleLib.TaskViewTasksFormat,
						row["recordID"].(int64),
						html.EscapeString(row["getMethod"].(string)),
						html.EscapeString(row["getIP"].(string)),
						html.EscapeString(row["getTime"].(string)),
						row["recordID"].(int64),
						row["recordID"].(int64),
						row["recordID"].(int64),
						row["recordID"].(int64),

					)

				}

				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"html":viewRow,
						"page":strconv.Itoa(currentPage + 1),
						"total": strconv.Itoa(pagesTotal),
					},
				)

			}

		}




	case "back":

		viewRow := ""
		pagesTotal := 1

		// get total
		r,_ := thisSql.Query(
			"appTaskView",
			[]string{"*"},
			[]string{"taskID","userID"},
			[]string{taskID,userID},
			" AND ",
			"",
		)

		total := len(r)
		showTotal := appConfig.RecordShowPerPage
		if total > 0 {

			pagesTotal = int(math.Ceil(float64(total) / float64(appConfig.RecordShowPerPage)))
			if total < appConfig.RecordShowPerPage {
				showTotal = total
			}

			for _, row := range r[:showTotal] {

				viewRow += fmt.Sprintf(moduleLib.TaskViewTasksFormat,
					row["recordID"].(int64),
					html.EscapeString(row["getMethod"].(string)),
					html.EscapeString(row["getIP"].(string)),
					html.EscapeString(row["getTime"].(string)),
					row["recordID"].(int64),
					row["recordID"].(int64),
					row["recordID"].(int64),
					row["recordID"].(int64),

				)

			}
		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"html":viewRow,
				"page": strconv.Itoa(1),
				"total":strconv.Itoa(pagesTotal),
			},
		)


	case "find":

		recordRow := ""
		recordID := c.DefaultPostForm("findRecordID","")



		// get total
		r,_ := thisSql.Query(
			"appTaskView",
			[]string{"*"},
			[]string{"recordID","taskID","userID"},
			[]string{recordID,taskID,userID},
			" AND ",
			"",
		)


		if len(r) == 1{


			recordRow += fmt.Sprintf(moduleLib.TaskViewTasksFormat,
				r[0]["recordID"].(int64),
				html.EscapeString(r[0]["getMethod"].(string)),
				html.EscapeString(r[0]["getIP"].(string)),
				html.EscapeString(r[0]["getTime"].(string)),
				r[0]["recordID"].(int64),
				r[0]["recordID"].(int64),
				r[0]["recordID"].(int64),
				r[0]["recordID"].(int64),

			)

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"ok",
					"html":recordRow,
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

	case "delete":

		recordID := c.DefaultPostForm("recordID","")
		updateCount := 0


		// delete

		r,_ := thisSql.Delete(
				"appTaskView",
				[]string{"recordID","taskID","userID"},
				[]string{recordID,taskID,userID},
				" AND ",
				"",
			)


		if r != 1{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Failed to delete record",
				},
			)

			return

		}

		// Get count

		rt,_ := thisSql.Query(
			"appTaskView",
			[]string{"*"},
			[]string{"taskID","userID"},
			[]string{taskID,userID},
			" AND ",
			"",
		)

		updateCount = len(rt)

		// update

		rts,_ := thisSql.Update(
			"appTasks",
			[]string{"taskRecordNum"},
			[]string{strconv.Itoa(updateCount)},
			[]string{"taskID","taskUserID"},
			[]string{taskID,userID},
			" AND ",
			"",
		)

		if rts != 1{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Failed to update task record",
				},
			)

			return

		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"refresh",
				"contents":"Record deleted successfully~",
			},
		)

	case "save":

		taskName := c.DefaultPostForm("taskName","")
		taskData := c.DefaultPostForm("taskData","")



		// check

		if !lib.CheckTaskName(taskName){
			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"The task name is not valid",
				},
			)
			return
		}

		// update

		r, _ := thisSql.Update(
				"appTasks",
				[]string{"taskName","taskData"},
				[]string{taskName,taskData},
				[]string{"taskID","taskUserID"},
				[]string{taskID,userID},
				" AND ",
				"",
			)

		if r != 1{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"warning",
					"contents":"The task data may not have changed",
				},
			)
			return

		}

		c.JSON(
			http.StatusOK,
			map[string]string{
				"status":"ok",
				"contents":"Task data changed successfully~",
				"taskName":taskName,
				"taskData":taskData,
			},
		)

	case "view":

		recordID := c.DefaultPostForm("recordID","")

		// Get record data
		r,_ := thisSql.Query(
				"appTaskView",
				[]string{"getResult","getModuleResult"},
				[]string{"recordID","taskID","userID"},
				[]string{recordID,taskID,userID},
				" AND ",
				"",
			)

		if len(r) != 1{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Unable to retrieve record",
				},
			)
			return

		}

		// Get module type

		rt,_ := thisSql.Query(
				"appTasks",
				[]string{"taskModuleID","taskModulePublic"},
				[]string{"taskID","taskUserID"},
				[]string{taskID,userID},
				" AND ",
				"",
			)

		if len(rt) != 1{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"error",
					"contents":"Unable to get module type",
				},
			)
			return

		}

		if rt[0]["taskModulePublic"].(int64) == 1{

			// Get module data
			rts,_ := thisSql.Query(
					"appModule",
					[]string{"*"},
					[]string{"moduleID"},
					[]string{strconv.FormatInt(rt[0]["taskModuleID"].(int64),10)},
					" AND ",
					"",
					)

			if len(rts) != 1{
				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"error",
						"contents":"Unable to get module data",
					},
				)
				return
			}

			switch rts[0]["moduleName"].(string){

			// default
			case publicModule.AllPublicModuleName["default"]:

				extraData := make(map[string]string)
				extraString := ""

				err := json.Unmarshal([]byte(r[0]["getModuleResult"].(string)),&extraData)
				if err != nil{

					c.JSON(
						http.StatusOK,
						map[string]string{
							"status":"error",
							"contents":"Unable to get extra result data",
						},
					)
					return

				}

				for name,value := range extraData{
					if len(value) > 0 && value!="\"\""{
						extraString += fmt.Sprintf(publicModule.Module_default_option, html.EscapeString(name), len(value)/64, html.EscapeString(value))
					}

				}

				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"returnType":"public_default",
						"reContents":r[0]["getResult"].(string),
						"extraContents":extraString,
					},
				)

			// request
			case publicModule.AllPublicModuleName["request"]:

				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"returnType":"public_request",
						"reContents":r[0]["getResult"].(string),
					},
				)

				// fish
			case publicModule.AllPublicModuleName["fish"]:

				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"returnType":"public_fish",
						"reContents":r[0]["getResult"].(string),
					},
				)

				// port
			case publicModule.AllPublicModuleName["port"]:

				c.JSON(
					http.StatusOK,
					map[string]string{
						"status":"ok",
						"returnType":"public_port",
						"reContents":r[0]["getResult"].(string),
						"extraContents":r[0]["getModuleResult"].(string),
					},
				)

			}




		}else{

			c.JSON(
				http.StatusOK,
				map[string]string{
					"status":"ok",
					"returnType":"custom",
					"reContents":r[0]["getResult"].(string),
				},
			)

		}

	}
	
}