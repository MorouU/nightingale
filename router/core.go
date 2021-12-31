package appRouter

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"net/http"
	webConfig "nightingale/config"
	"nightingale/lib/globals"
)

func RouterInit(){
	// 设置发行模式
	gin.SetMode(gin.ReleaseMode)
	// 获取Gin框架主路由
	thisRouter := gin.Default()
	// 设置模板页面所在的路径
	thisRouter.LoadHTMLGlob(webConfig.AppPathConfig["templatePath"])
	// 设置静态资源页面所在的路径
	thisRouter.StaticFS("/static",http.Dir(webConfig.AppPathConfig["staticPath"]))
	// 将当前路由变量存入全局映射
	globals.Set("routerThis",thisRouter,true)
	// 设置根页面路由，并存入全局映射
	globals.Set("routerIndex",thisRouter.Group("/"),true)
	// 设置用户登录功能路由，并存入全局映射
	globals.Set("routerUserInterfaces",thisRouter.Group("/userEnter"),true)
	// 设置用户主页功能路由，并存入全局映射
	globals.Set("routerUserProfile",thisRouter.Group("/userProfile"),true)
	// 设置API功能路由，并存入全局映射
	globals.Set("routerApi",thisRouter.Group("/api"),true)
	// 设置管理员登录功能路由，并存入全局映射
	globals.Set("routerAdminInterfaces",thisRouter.Group("/adminEnter"),true)
	// 设置管理员主页功能路由，并存入全局映射
	globals.Set("routerAdminProfile",thisRouter.Group("/adminProfile"),true)
	// 初始化session中间件
	sessionInit()
	// 初始化页面路由所指向函数
	pagesInit()
	// 在80端口开启当前服务器
	thisRouter.Run(":80")
}


func sessionInit(){
	// 从全局映射中取出各个路由
	routerIndex := globals.Get("routerIndex").(*gin.RouterGroup)
	routerUserInterfaces := globals.Get("routerUserInterfaces").(*gin.RouterGroup)
	routerUserProfile := globals.Get("routerUserProfile").(*gin.RouterGroup)
	routerAdminInterfaces := globals.Get("routerAdminInterfaces").(*gin.RouterGroup)
	routerAdminProfile := globals.Get("routerAdminProfile").(*gin.RouterGroup)
	thisRouter := globals.Get("routerUserInterfaces").(*gin.RouterGroup)
	// 设置session密钥
	serectKey := webConfig.SessionConfig["serectKey"].(string)
	if serectKey == ""{
		serectKey = uuid.New()
	}
	store := sessions.NewCookieStore([]byte(serectKey))
	// 设置session配置
	store.Options(sessions.Options{
		Path:     webConfig.SessionConfig["path"].(string),
		Domain:   webConfig.SessionConfig["domain"].(string),
		MaxAge:   webConfig.SessionConfig["maxAge"].(int),
		Secure:   webConfig.SessionConfig["secure"].(bool),
		HttpOnly: webConfig.SessionConfig["httpOnly"].(bool),
	})
	// 为每一个路由应用session配置
	routerIndex.Use(sessions.Sessions(webConfig.SessionConfig["cookieName"].(string),store))
	routerUserInterfaces.Use(sessions.Sessions(webConfig.SessionConfig["cookieName"].(string),store))
	routerUserProfile.Use(sessions.Sessions(webConfig.SessionConfig["cookieName"].(string),store))
	routerAdminInterfaces.Use(sessions.Sessions(webConfig.SessionConfig["cookieName"].(string),store))
	routerAdminProfile.Use(sessions.Sessions(webConfig.SessionConfig["cookieName"].(string),store))
	thisRouter.Use(sessions.Sessions(webConfig.SessionConfig["cookieName"].(string),store))
}

func pagesInit(){
	// 从全局映射中取出各个路由，并设置对应路由所指向函数
	routerIndex := globals.Get("routerIndex").(*gin.RouterGroup)
	routerIndex.Any("", pages_index_index)
	routerUserInterfaces := globals.Get("routerUserInterfaces").(*gin.RouterGroup)
	routerUserInterfaces.Any("/login", pages_userEnter_login)
	routerUserInterfaces.Any("/register", pages_userEnter_reg)
	routerUserInterfaces.Any("/forget", pages_userEnter_getpass)
	routerUserInterfaces.POST("/authCodeRefresh", api_userEnter_loginAuthCodeRefresh)
	routerUserInterfaces.GET("/api/userActive/:activeCode", api_userEnter_registerActive)

	routerUserProfile := globals.Get("routerUserProfile").(*gin.RouterGroup)
	routerUserProfile.Any("/:uuid/profile", pages_userProfile_profile)
	routerUserProfile.Any("/:uuid/modules", pages_userProfile_modules)
	routerUserProfile.Any("/:uuid/home", pages_userProfile_home)
	routerUserProfile.Any("/:uuid/tasks", pages_userProfile_tasks)
	routerUserProfile.Any("/:uuid/:taskID/view", pages_userProfile_view)
	routerUserProfile.Any("/:uuid/logout", pages_userProfile_logout)

	routerAdminInterfaces := globals.Get("routerAdminInterfaces").(*gin.RouterGroup)
	routerAdminInterfaces.Any("/login", pages_adminEnter_login)
	routerAdminInterfaces.POST("/authCodeRefresh", api_adminEnter_loginAuthCodeRefresh)

	routerAdminProfile := globals.Get("routerAdminProfile").(*gin.RouterGroup)
	routerAdminProfile.Any("/:uuid/profile", pages_adminProfile_profile)
	routerAdminProfile.Any("/:uuid/userList", pages_adminProfile_userList)
	routerAdminProfile.Any("/:uuid/:userID/tasks", pages_adminProfile_userTasks)
	routerAdminProfile.Any("/:uuid/:userID/modules", pages_adminProfile_userModules)
	routerAdminProfile.Any("/:uuid/logout", pages_adminProfile_logout)

	routerApi := globals.Get("routerApi").(*gin.RouterGroup)
	routerApi.Any("/:taskApi", pages_api_firstLink)
	routerApi.Any("/:taskApi/default/:randStr", pages_api_secondLink)
	routerApi.Any("/:taskApi/port/:randStr", pages_api_portLink)
}