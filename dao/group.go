package dao

import (
	"FIFA/model"
	"errors"
	"gorm.io/gorm"
)

func GetGroupById(id int) (group *model.Group, err error) {
	err = DB.Model(&model.Group{}).Where("id = ?", id).First(&group).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("group is not exist")
	}
	return
}

func GetGroupList() (groups []*model.Group, count int64, err error) {
	err = DB.Model(&model.Group{}).Count(&count).Find(&groups).Error
	return
}
