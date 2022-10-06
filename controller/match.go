package controller

import (
	"FIFA/controller/response"
	"FIFA/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetMatchHandler
// @Description 查找单次比赛
// @Summary 查询单次比赛接口
// @Tags Match
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id query int true "比赛id"
// @Success 200 {object} response.ResponseData
// @Router /match/:id [get]
func GetMatchHandler(c *gin.Context) {
	matchId, _ := strconv.Atoi(c.Param("id"))
	match, err := service.GetMatchById(matchId)
	if err != nil {
		zap.L().Error("service GetMatchById failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, match)
}

// GetMatchListHandler
//@Description 查找比赛
//@Summary 查询所有比赛接口
// @Tags Match
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.ResponseData
//@Router /match [get]
func GetMatchListHandler(c *gin.Context) {
	matches, count, err := service.GetMatchList()
	if err != nil {
		zap.L().Error("service GetMatchList failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count": count,
		"Match": matches,
	})
}
