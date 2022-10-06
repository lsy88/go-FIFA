package controller

import (
	"FIFA/controller/response"
	"FIFA/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetCoachListHandler
//@Description 查找教练
//@Summary 查询所有教练接口
// @Tags Coach
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.ResponseData
//@Router /coach [get]
func GetCoachListHandler(c *gin.Context) {
	coaches, count, err := service.GetCoachList()
	if err != nil {
		zap.L().Error("service GetCoachList failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count": count,
		"Coach": coaches,
	})
}

// GetCoachHandler
//@Description 查找单教练
//@Summary 查询单个教练接口
// @Tags Coach
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id query int true "教练ID"
// @Success 200 {object} response.ResponseData
//@Router /coach/:id [get]
func GetCoachHandler(c *gin.Context) {
	coachId, _ := strconv.Atoi(c.Param("id"))
	coach, err := service.GetCoachById(coachId)
	if err != nil {
		zap.L().Error("service GetCoachById failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, coach)
}
