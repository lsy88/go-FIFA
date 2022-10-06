package router

import (
	"FIFA/controller"
	"github.com/gin-gonic/gin"
)

func ClassicRouter(v1 *gin.RouterGroup) {
	v1.GET("/classic/:year", controller.GetClassicListByYearHandler)
	v1.GET("/classic/list", controller.GetClassicListHandler)
}
