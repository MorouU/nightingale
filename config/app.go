package appConfig

var Domain = "http://xss.ni/"
var ApiDomain = "http://xss.ni/api/%s"
var CookieDomain = "xss.ni"

var TaskShowPerPage = 6

var RecordShowPerPage = 6

var UserLevelString = map[int64]string{

	0:"低级用户",
	1:"普通用户",
	2:"高级用户",

}

var UserLevelStringFromString = map[string]string{

	"0":"低级用户",
	"1":"普通用户",
	"2":"高级用户",

}

var AppPathConfig =map[string]string{
	"authCodeFile":"E:/g/NGF/src/nightingale/config/Bad Comic.ttf",
	"templatePath":"E:/g/NGF/src/nightingale/router/pages/**/*",
	"staticPath":"E:/g/NGF/src/nightingale/static",
	"uploadPath":"E:/g/NGF/src/nightingale/upload",
}

var UserListShowPerPage = 8

var TaskListShowPerPage = 8

var ModuleListShowPerPage = 8