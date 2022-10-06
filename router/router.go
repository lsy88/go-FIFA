package router

import (
	"FIFA/core"
	_ "FIFA/docs"
	"FIFA/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"                  // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"net/http"
)

func SetupRouter(mode string) {
	if mode == "dev" {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	v1 := r.Group("/FIFA/v1/api")
	//健康检测
	v1.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"Welcome": "FIFA-World-Cup",
		})
	})
	
	AdminRouter(v1)
	v1.Use(middleware.AuthCheck()) //使用jwt验证
	{
		PlayerRouter(v1)
		AwardPlayerRouter(v1)
		TeamRouter(v1)
		MatchRouter(v1)
		GroupRouter(v1)
		CoachRouter(v1)
		AwardRouter(v1)
		ClassicRouter(v1)
		StatisticsRouter(v1)
	}
	
	r.Run(":" + core.Conf.Port)
}
