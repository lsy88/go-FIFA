package controller

import (
	"FIFA/controller/request"
	"FIFA/controller/response"
	"FIFA/service"
	"FIFA/utils/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AdminLoginHandler 管理员登录
// @Summary 登录业务接口
// @Description 登录业务接口
// @Tags Admin
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "姓名"
// @Param password formData string true "密码"
// @Success 200 {object} response.ResponseData
// @Router /admin/login [post]
func AdminLoginHandler(c *gin.Context) {
	var user request.ReqAdminLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ResponseError(c, response.ERR_Invalid_Params)
		return
	}
	if err := service.AdminLogin(&user); err != nil {
		zap.L().Error("service Login failed", zap.Error(err))
		response.ResponseError(c, response.ERR_NotLogin)
		return
	}
	//生成jwt
	j := jwt.NewJWT()
	token, _ := j.GenToken(user.Name)
	response.ResponseSuccess(c, gin.H{
		"accessToken": token,
		"name":        user.Name,
	})
}

// AdminRegisterHandler 管理员注册
// @Summary 管理员注册接口
// @Description 管理员注册接口
// @Tags Admin
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "姓名"
// @Param password formData string true "密码"
// @Param phone formData string false "电话"
// @Security Bearer
// @Success 200 {object} response.ResponseData
// @Router /admin/register [POST]
func AdminRegisterHandler(c *gin.Context) {
	var user request.ReqAdminRegister
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ResponseError(c, response.ERR_Invalid_Params)
		return
	}
	if err := service.AdminRegister(&user); err != nil {
		zap.L().Error("service register failed", zap.Error(err))
		response.ResponseError(c, response.ERR_NotRegister)
		return
	}
	response.ResponseSuccess(c, "注册成功")
}
