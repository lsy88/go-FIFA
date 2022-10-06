package router

import (
	"FIFA/controller"
	"github.com/gin-gonic/gin"
)

func AwardRouter(v1 *gin.RouterGroup) {
	v1.GET("/award", controller.GetAwardListHandler)
	v1.GET("/award/:id", controller.GetAwardHandler)
	//根据关键字查询奖项
	v1.GET("/award/info/:keyword", controller.GetAwardByKeyHandler)
}
