package controller

import (
	"FIFA/controller/response"
	"FIFA/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetAwardPlayerHandler
//@Description 查找奖牌获得者
//@Summary 查询所有奖牌获得者接口
// @Tags AwardPlayer
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.ResponseData
//@Router /award_player [get]
func GetAwardPlayerHandler(c *gin.Context) {
	awardPlayer, count, err := service.GetAwardPlayerList()
	if err != nil {
		zap.L().Error("service GetAwardPlayerList failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count":       count,
		"AwardPlayer": awardPlayer,
	})
}

// GetAwardPlayerByKeyHandler
//@Description 根据关键字查找奖牌得主
//@Summary 查询奖牌得主接口
// @Tags AwardPlayer
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param keyword query string true "关键字"
// @Success 200 {object} response.ResponseData
//@Router /award_player/:keyword [get]
func GetAwardPlayerByKeyHandler(c *gin.Context) {
	keyword := c.Param("keyword")
	awardPlayers, count, err := service.GetAwardPlayerByKeyword(keyword)
	if err != nil {
		zap.L().Error("service GetAwardPlayerByKeyword failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count":        count,
		"awardPlayers": awardPlayers,
	})
}
