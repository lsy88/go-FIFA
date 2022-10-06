package dao

import (
	"FIFA/controller/request"
	"FIFA/model"
	"errors"
)

func AdminLogin(user *request.ReqAdminLogin) (err error) {
	var admin model.Admin
	err = DB.Where("name = ?", user.Name).First(&admin).Error
	if err != nil {
		return
	}
	if !admin.PasswordCheck(user.Password) {
		return errors.New("Err_Admin_Password")
	}
	return nil
}

func AdminRegister(user *model.Admin) (err error) {
	err = DB.Model(&model.Admin{}).Create(&user).Error
	return
}
