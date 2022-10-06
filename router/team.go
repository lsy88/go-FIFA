package router

import (
	"FIFA/controller"
	"github.com/gin-gonic/gin"
)

func TeamRouter(v1 *gin.RouterGroup) {
	v1.GET("/team", controller.GetTeamListHandler)
	v1.GET("/team/:id", controller.GetTeamHandler)
}
