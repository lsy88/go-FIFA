package dao

import (
	"FIFA/model"
	"errors"
	"gorm.io/gorm"
)

func GetMatchById(id int) (match *model.Match, err error) {
	err = DB.Model(&model.Match{}).Where("id = ?", id).First(&match).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("match is not exist")
	}
	return
}

func GetMatchList() (matches []*model.Match, count int64, err error) {
	err = DB.Model(&model.Match{}).Count(&count).Find(&matches).Error
	return
}
