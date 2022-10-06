package router

import (
	"FIFA/controller"
	"github.com/gin-gonic/gin"
)

func CoachRouter(v1 *gin.RouterGroup) {
	v1.GET("/coach", controller.GetCoachListHandler)
	v1.GET("/coach/:id", controller.GetCoachHandler)
}
