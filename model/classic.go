package model

import (
	"gorm.io/gorm"
)

//往届赛事存档
type WorldCupArchive struct {
	gorm.Model
	URL         string `json:"url" gorm:"type:varchar(128);column:world_cup_url"`
	Name        string `json:"name" gorm:"type:varchar(128);column:country_name"`
	Year        string `json:"year" gorm:"type:varchar(64);column:year"`
	Winner      string `json:"winner" gorm:"type:varchar(64);column:winner_country"`      //冠军
	RunnersUp   string `json:"runners_up" gorm:"type:varChar(64);column:runners_up_name"` //亚军
	Third       string `json:"third_name" gorm:"type:varchar(64);column:third_name"`      //第三名
	Fourth      string `json:"fourth_name" gorm:"type:varchar(64);column:fourth_name"`    //第四名
	FinalResult string `json:"final_result" gorm:"type:varchar(128);column:final_result"`
	Title       string `json:"title" gorm:"type:varchar(64);column:title"`
}
