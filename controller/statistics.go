package controller

import (
	"FIFA/controller/request"
	"FIFA/controller/response"
	"FIFA/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetPlayerGoalByRankHandler
//@Description 查找排名
//@Summary 查询排名接口
// @Tags Statistics
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param rank query string true "排名"
// @Success 200 {object} response.ResponseData
//@Router /statistics/players/goals/:rank [get]
//根据排名查询信息
func GetPlayerGoalByRankHandler(c *gin.Context) {
	rank, _ := strconv.Atoi(c.Param("rank"))
	player, err := service.GetPlayerGoalByRank(rank)
	if err != nil {
		zap.L().Error("service GetPlayerGoalByRank failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, player)
}

// GetPlayerGoalHandler
//@Description 查找球员奖牌
//@Summary 查询球员得奖接口
// @Tags Statistics
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param name  formData string false "姓名"
// @Param goal formData string false "得分"
// @Success 200 {object} response.ResponseData
//@Router /statistics/players/goals [get]
func GetPlayerGoalHandler(c *gin.Context) {
	var request request.ReqPlayerGoal
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.L().Error("ShouldBindJSON ReqPlayerGoal failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Invalid_Params)
		return
	}
	players, count, err := service.GetPlayerGoalHandler(&request)
	if err != nil {
		zap.L().Error("service GetPlayerGoalHandler failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count":   count,
		"Players": players,
	})
}

// GetSaveByRankHandler
//@Description 查找排名
//@Summary 查询排名接口
// @Tags Statistics
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param rank query string true "排名"
// @Success 200 {object} response.ResponseData
//@Router /statistics/players/saves/:rank [get]
func GetSaveByRankHandler(c *gin.Context) {
	rank, _ := strconv.Atoi(c.Param("rank"))
	player, err := service.GetPlayerSaveByRank(rank)
	if err != nil {
		zap.L().Error("service GetPlayerSavaByRank failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, player)
}

// GetPlayerSaveHandler
//@Description 查找球员奖牌
//@Summary 查询球员得奖接口
// @Tags Statistics
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param name  query string false  "排名"
// @Success 200 {object} response.ResponseData
//@Router /statistics/players/saves [get]
func GetPlayerSaveHandler(c *gin.Context) {
	name := c.Query("name")
	players, count, err := service.GetPlayerSave(name)
	if err != nil {
		zap.L().Error("service GetPlayerSave failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count":   count,
		"Players": players,
	})
}

func GetShotByRankHandler(c *gin.Context) {
	rank, _ := strconv.Atoi(c.Param("rank"))
	player, err := service.GetPlayerShotByRank(rank)
	if err != nil {
		zap.L().Error("service GetPlayerShotByRank failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, player)
}

func GetPlayerShotHandler(c *gin.Context) {
	name := c.Query("name")
	players, count, err := service.GetPlayerShot(name)
	if err != nil {
		zap.L().Error("service GetPlayerShot failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count":   count,
		"Players": players,
	})
}

func GetPlayerDisciplinaryByRankHandler(c *gin.Context) {
	rank, _ := strconv.Atoi(c.Param("rank"))
	player, err := service.GetPlayerDisciplinaryByRank(rank)
	if err != nil {
		zap.L().Error("service GetPlayerDisciplinaryByRank failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, player)
}

func GetPlayerDisciplinaryHandler(c *gin.Context) {
	players, count, err := service.GetPlayerDisciplinary()
	if err != nil {
		zap.L().Error("service GetPlayerDisciplinary failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count":   count,
		"Players": players,
	})
}

func GetTeamGoalByRankHandler(c *gin.Context) {
	rank, _ := strconv.Atoi(c.Param("rank"))
	team, err := service.GetTeamGoalByRank(rank)
	if err != nil {
		zap.L().Error("service GetTeamGoalByRank failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, team)
}

func GetTeamGoalHandler(c *gin.Context) {
	teams, count, err := service.GetTeamGoal()
	if err != nil {
		zap.L().Error("service GetTeamGoal failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count": count,
		"Teams": teams,
	})
}

func GetTeamAttemptsByRankHandler(c *gin.Context) {
	rank, _ := strconv.Atoi(c.Param("rank"))
	team, err := service.GetTeamAttemptsByRank(rank)
	if err != nil {
		zap.L().Error("service GetTeamAttemptsByRank failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, team)
}

func GetTeamAttemptsHandler(c *gin.Context) {
	teams, count, err := service.GetTeamAttempts()
	if err != nil {
		zap.L().Error("service GetTeamAttempts failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count": count,
		"Teams": teams,
	})
}

func GetTeamDisciplinaryByRankHandler(c *gin.Context) {
	rank, _ := strconv.Atoi(c.Param("rank"))
	team, err := service.GetTeamDisciplinaryByRank(rank)
	if err != nil {
		zap.L().Error("service GetTeamDisciplinaryByRank failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, team)
}

func GetTeamDisciplinaryHandler(c *gin.Context) {
	teams, count, err := service.GetTeamDisciplinary()
	if err != nil {
		zap.L().Error("service GetTeamDisciplinary failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count": count,
		"Teams": teams,
	})
}
