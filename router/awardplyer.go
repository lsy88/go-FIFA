package router

import (
	"FIFA/controller"
	"github.com/gin-gonic/gin"
)

func AwardPlayerRouter(v1 *gin.RouterGroup) {
	v1.GET("/award_player", controller.GetAwardPlayerHandler)
	v1.GET("/award_player/:keyword", controller.GetAwardPlayerByKeyHandler)
}
