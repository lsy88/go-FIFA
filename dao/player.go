package dao

import (
	"FIFA/model"
	"errors"
	"gorm.io/gorm"
)

func GetPlayerById(id int) (player *model.Player, err error) {
	err = DB.Model(&model.Player{}).Where("id = ?", id).Omit("CreatedAt", "UpdatedAt", "DeletedAt").First(&player).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("player is not exist")
	}
	return
}

func GetPlayerList() (players []*model.Player, count int64, err error) {
	err = DB.Model(&model.Player{}).Count(&count).Omit("CreatedAt", "UpdatedAt", "DeletedAt").Find(&players).Error
	return
}
