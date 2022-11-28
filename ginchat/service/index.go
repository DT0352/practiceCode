package service

import (
	"chat/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, "我饿了")
}
func GetUserList(c *gin.Context) {

	c.JSON(http.StatusOK, model.GetUserList())
}
