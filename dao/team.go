package dao

import (
	"FIFA/model"
	"errors"
	"gorm.io/gorm"
)

func GetTeamById(id int) (team *model.Team, err error) {
	err = DB.Model(&model.Team{}).Where("id = ?", id).First(&team).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("team is not exist")
	}
	return
}

func GetTeamList() (teams []*model.Team, count int64, err error) {
	err = DB.Model(&model.Team{}).Count(&count).Find(&teams).Error
	return
}
