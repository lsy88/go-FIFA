package controller

import (
	"FIFA/controller/response"
	"FIFA/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetGroupListHandler
//@Description 查找小组
//@Summary 查询所有小组接口
// @Tags Group
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.ResponseData
//@Router /group [get]
func GetGroupListHandler(c *gin.Context) {
	groups, count, err := service.GetGroupList()
	if err != nil {
		zap.L().Error("service GetGroupList failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, gin.H{
		"Count": count,
		"Group": groups,
	})
}

// GetGroupHandler
//@Description 查找单个小组
//@Summary 查询单个小组接口
// @Tags Group
//@Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id query int true "小组id"
// @Success 200 {object} response.ResponseData
//@Router /group/:id [get]
func GetGroupHandler(c *gin.Context) {
	groupId, _ := strconv.Atoi(c.Param("id"))
	group, err := service.GetGroupById(groupId)
	if err != nil {
		zap.L().Error("service GetGroupById failed", zap.Error(err))
		response.ResponseError(c, response.ERR_Server_Busy)
		return
	}
	response.ResponseSuccess(c, group)
}
