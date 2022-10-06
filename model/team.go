package model

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Flag string `json:"flag" gorm:"type:varchar(128); not null; column:flag_address"`
	Name string `json:"name" gorm:"type:varchar(64); not null; column:team_name"`
}
