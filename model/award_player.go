package model

import "gorm.io/gorm"

type AwardPlayer struct {
	gorm.Model
	AwardName  string `json:"award_name" gorm:"award_name"`
	PlayerName string `json:"player_name" gorm:"player_name"`
	Country    string `json:"country" gorm:"country"`
}
