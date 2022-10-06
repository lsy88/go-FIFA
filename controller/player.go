package controller

import (
	"FIFA/controller/response"
	"FIFA/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetPlayerListHandler
//@Description 查找球员
//@Summary 查询所有球员接口
// @Tags Player
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.ResponseData
//@Router /player [get]
func GetPlayerListHandler(c *gin.Context) {
	players, count, err := service.GetPlayerList()
	if err != nil {
		zap.L().Error("service GetPlayerList failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count":  count,
		"Player": players,
	})
}

// GetPlayerHandler
//@Description 查找单个球员
//@Summary 查询单个球员接口
// @Tags Player
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id query int true "球员id"
// @Success 200 {object} response.ResponseData
//@Router /player/:id [get]
func GetPlayerHandler(c *gin.Context) {
	playerId, _ := strconv.Atoi(c.Param("id"))
	player, err := service.GetPlayerById(playerId)
	if err != nil {
		zap.L().Error("service GetPlayerById failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, player)
}
