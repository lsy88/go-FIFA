package dao

import (
	"FIFA/model"
	"errors"
	"gorm.io/gorm"
)

func GetClassicByYear(year int) (classic *model.WorldCupArchive, err error) {
	err = DB.Model(&model.WorldCupArchive{}).Where("year = ?", year).First(&classic).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("不存在该年份比赛")
	}
	return
}

func GetClassicAll() (classics []*model.WorldCupArchive, count int64, err error) {
	err = DB.Model(&model.WorldCupArchive{}).Count(&count).Find(&classics).Error
	return
}

func GetClassicByTitle(title string) (classics []*model.WorldCupArchive, count int64, err error) {
	err = DB.Model(&model.WorldCupArchive{}).Where("title like ?", "%"+title+"%").
		Count(&count).Find(&classics).Error
	return
}

func GetClassicByResult(result string) (classics []*model.WorldCupArchive, count int64, err error) {
	err = DB.Model(&model.WorldCupArchive{}).Where("final_result like ?", "%"+result+"%").
		Count(&count).Find(&classics).Error
	return
}
