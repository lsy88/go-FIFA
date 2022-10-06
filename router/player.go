package router

import (
	"FIFA/controller"
	"github.com/gin-gonic/gin"
)

func PlayerRouter(v1 *gin.RouterGroup) {
	v1.GET("/player", controller.GetPlayerListHandler)
	v1.GET("/player/:id", controller.GetPlayerHandler)
}
