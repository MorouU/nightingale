package appRouter

import (
	"github.com/gin-gonic/gin"
	lib "nightingale/lib/func"
	"net/http"
)

func pages_index_index(c *gin.Context){

	reqIP := c.ClientIP()
	reqTime := lib.GetTimeDateTime()

	sessionSet(c, map[interface{}]interface{}{
		"reqIP":reqIP,
		"reqTime":reqTime,
	})


	c.HTML(
		http.StatusOK,
		"index_index.ni",
		gin.H{},
		)
}
