package router

import (
	"chat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("index", service.Index)
	r.GET("userList", service.GetUserList)
	return r
}
