package router

import (
	"FIFA/controller"
	"github.com/gin-gonic/gin"
)

func MatchRouter(v1 *gin.RouterGroup) {
	v1.GET("/match", controller.GetMatchListHandler)
	v1.GET("/match/:id", controller.GetMatchHandler)
}
