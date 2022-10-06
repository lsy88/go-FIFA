package controller

import (
	"FIFA/controller/response"
	"FIFA/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetAwardListHandler
//@Description 查找奖牌
//@Summary 查询所有奖牌接口
// @Tags Award
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.ResponseData
//@Router /award [get]
func GetAwardListHandler(c *gin.Context) {
	awards, count, err := service.GetAwardList()
	if err != nil {
		zap.L().Error("service GetAwardList failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count":  count,
		"Awards": awards,
	})
}

// GetAwardHandler
//@Description 查找单奖牌
//@Summary 查询单个奖牌接口
// @Tags Award
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id query int true "奖牌id"
// @Success 200 {object} response.ResponseData
//@Router /award/:id [get]
func GetAwardHandler(c *gin.Context) {
	awardId, _ := strconv.Atoi(c.Param("id"))
	award, err := service.GetAwardById(awardId)
	if err != nil {
		zap.L().Error("service GetAwardById failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, award)
}

// GetAwardByKeyHandler
//@Description 根据关键字查找奖牌
//@Summary 查询单个奖牌接口
// @Tags Award
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param keyword query string true "关键字"
// @Success 200 {object} response.ResponseData
//@Router /award/info/:keyword [get]
func GetAwardByKeyHandler(c *gin.Context) {
	keyword := c.Param("keyword")
	awards, count, err := service.GetAwardByKeyword(keyword)
	if err != nil {
		zap.L().Error("service GetAwardByKeyword failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count":  count,
		"Awards": awards,
	})
}
