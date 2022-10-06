package router

import (
	"FIFA/controller"
	"github.com/gin-gonic/gin"
)

func AdminRouter(v1 *gin.RouterGroup) {
	v1.POST("/admin/login", controller.AdminLoginHandler)
	v1.POST("/admin/register", controller.AdminRegisterHandler)
}
