package dao

import (
	"FIFA/model"
	"errors"
	"gorm.io/gorm"
)

func GetAwardById(id int) (award *model.Award, err error) {
	err = DB.Model(&model.Award{}).Where("id = ?", id).First(&award).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("award is not exist")
	}
	return
}

func GetAwardList() (awards []*model.Award, count int64, err error) {
	err = DB.Model(&model.Award{}).Count(&count).Find(&awards).Error
	return
}

// GetAwardByKeyword 根据关键字查找奖项
func GetAwardByKeyword(keyword string) (awards []*model.Award, count int64, err error) {
	err = DB.Model(&model.Award{}).Where("info like ?", "%"+keyword+"%").
		Count(&count).Find(&awards).Error
	return
}
