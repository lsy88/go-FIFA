package controller

import (
	"FIFA/controller/request"
	"FIFA/controller/response"
	"FIFA/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetClassicListByYearHandler
//@Description 根据年份查找记录
//@Summary 查询比赛接口
// @Tags Classic
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param year query string true  "年份"
// @Success 200 {object} response.ResponseData
//@Router /classic/:year [get]
func GetClassicListByYearHandler(c *gin.Context) {
	year, _ := strconv.Atoi(c.Param("year"))
	classic, err := service.GetClassicByYear(year)
	if err != nil {
		zap.L().Error("service GetClassicByYear failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, classic)
}

// GetClassicListHandler
//@Description 查找比赛列表
//@Summary 查询所有比赛接口
// @Tags Classic
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.ResponseData
//@Router /classic/list [get]
func GetClassicListHandler(c *gin.Context) {
	var req request.ReqClassicMes
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("ShouldBindJSON ReqClassicMes failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Invalid_Params)
		return
	}
	classicList, count, err := service.GetClassicList(&req)
	if err != nil {
		zap.L().Error("service GetClassicList failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count": count,
		"List":  classicList,
	})
}
