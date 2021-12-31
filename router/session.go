package appRouter

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func sessionSet(c *gin.Context,sMap map[interface{}]interface{}){

	s := sessions.Default(c)

	for key,val := range sMap{
		s.Set(key,val)
	}
	s.Save()

}

func sessionGet(c *gin.Context,key interface{})(i interface{}){
	s:=sessions.Default(c)
	i = s.Get(key)
	s.Save()
	return

}

func sessionKeyExist(i interface{},key string) bool {
	switch key {
	case "int":
		_, val := i.(int)
		return val
	case "int64":
		_, val := i.(int64)
		return val
	case "string":
		_, val := i.(string)
		return val
	default:
		return false
	}

}

func sessionClear(c *gin.Context) bool {
	s := sessions.Default(c)

	if sessionKeyExist(sessionGet(c,"isLogin"),"string") {
		sessionSet(c, map[interface{}]interface{}{
			"isLogin":"0",
		})
	}

	if sessionKeyExist(sessionGet(c,"isAdmin"),"string") {
		sessionSet(c, map[interface{}]interface{}{
			"isAdmin":"0",
		})
	}

	s.Clear()
	return true
}