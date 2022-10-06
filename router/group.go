package router

import (
	"FIFA/controller"
	"github.com/gin-gonic/gin"
)

func GroupRouter(v1 *gin.RouterGroup) {
	v1.GET("/group", controller.GetGroupListHandler)
	v1.GET("/group/:id", controller.GetGroupHandler)
}
