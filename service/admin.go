package service

import (
	"FIFA/controller/request"
	"FIFA/dao"
	"FIFA/model"
	"FIFA/utils/check"
	"errors"
)

func AdminLogin(user *request.ReqAdminLogin) (err error) {
	if check.CheckName(user.Name) {
		err = dao.AdminLogin(user)
		if err != nil {
			return
		}
	} else {
		return errors.New("姓名格式有误")
	}
	return nil
}

func AdminRegister(user *request.ReqAdminRegister) (err error) {
	if !check.CheckName(user.Name) {
		return errors.New("姓名格式有误")
	}
	admin := model.Admin{
		Name:              user.Name,
		Phone:             user.Phone,
		EncryptedPassword: check.PassWordEncrypted(user.Password),
	}
	if err = dao.AdminRegister(&admin); err != nil {
		return err
	}
	return nil
}
