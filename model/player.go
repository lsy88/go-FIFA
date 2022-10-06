package model

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Number   string `json:"number" gorm:"type: varchar(12); not null; column:number"`
	Name     string `json:"name" gorm:"type: varchar(64); not null; column:name"`
	Country  string `json:"country" gorm:"type: varchar(64); not null; column:country"`
	Role     string `json:"role" gorm:"type: varchar(64); not null; column:role"`
	ImageURL string `json:"image_address" gorm:"type: varchar(128); not null; column:image_address"`
}

//教练
type Coach struct {
	gorm.Model
	CountryName string `json:"country_name" gorm:"type:varchar(64); not null; column:country_name"`
	Name        string `json:"name" gorm:"type:varchar(64); not null; column:name"`
	ImageURL    string `json:"image_url" gorm:"type:varchar(128); not null; column:image_address"`
}
