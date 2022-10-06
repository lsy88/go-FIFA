package controller

import (
	"FIFA/controller/response"
	"FIFA/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetTeamListHandler
//@Description 查找球队
//@Summary 查询所有球队接口
// @Tags Team
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.ResponseData
//@Router /team [get]
func GetTeamListHandler(c *gin.Context) {
	teams, count, err := service.GetTeamList()
	if err != nil {
		zap.L().Error("service GetTeamList failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count": count,
		"Teams": teams,
	})
}

// GetTeamHandler
//@Description 查找单个球队
//@Summary 查询单个球队接口
// @Tags Team
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id query int true "队伍id"
// @Success 200 {object} response.ResponseData
//@Router /team/:id [get]
func GetTeamHandler(c *gin.Context) {
	teamId, _ := strconv.Atoi(c.Param("id"))
	team, err := service.GetTeamById(teamId)
	if err != nil {
		zap.L().Error("service GetTeamById failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, team)
}
