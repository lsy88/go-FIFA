package router

import (
	"FIFA/controller"
	"github.com/gin-gonic/gin"
)

func StatisticsRouter(v1 *gin.RouterGroup) {
	{
		v1.GET("/statistics/players/goals/:rank", controller.GetPlayerGoalByRankHandler)
		v1.GET("/statistics/players/goals", controller.GetPlayerGoalHandler)
		
		v1.GET("/statistics/players/saves/:rank", controller.GetSaveByRankHandler)
		v1.GET("/statistics/players/saves", controller.GetPlayerSaveHandler)
		
		v1.GET("/statistics/players/shots/:rank", controller.GetShotByRankHandler)
		v1.GET("/statistics/players/shots", controller.GetPlayerShotHandler)
		
		v1.GET("/statistics/players/disciplinary/:rank", controller.GetPlayerDisciplinaryByRankHandler)
		v1.GET("/statistics/players/disciplinary", controller.GetPlayerDisciplinaryHandler)
	}
	{
		v1.GET("/statistics/teams/goals/:rank", controller.GetTeamGoalByRankHandler)
		v1.GET("/statistics/teams/goals", controller.GetTeamGoalHandler)
		
		v1.GET("/statistics/teams/attempts/:rank", controller.GetTeamAttemptsByRankHandler)
		v1.GET("/statistics/teams/attempts", controller.GetTeamAttemptsHandler)
		
		v1.GET("/statistics/teams/disciplinary/:rank", controller.GetTeamDisciplinaryByRankHandler)
		v1.GET("/statistics/teams/disciplinary", controller.GetTeamDisciplinaryHandler)
	}
}
