package dao

import (
	"FIFA/model"
	"errors"
	"gorm.io/gorm"
)

func GetCoachById(id int) (coach *model.Coach, err error) {
	err = DB.Model(&model.Coach{}).Where("id = ?", id).First(&coach).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("coach is not exist")
	}
	return
}

func GetCoachList() (coaches []*model.Coach, count int64, err error) {
	err = DB.Model(&model.Coach{}).Count(&count).Find(&coaches).Error
	return
}
