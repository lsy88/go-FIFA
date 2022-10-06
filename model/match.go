package model

import (
	"gorm.io/gorm"
	"time"
)

type Match struct {
	gorm.Model
	Date         string `json:"date" gorm:"type:varchar(64);not null;column:date"`
	GroupName    string `json:"group_name" gorm:"type:varchar(64);not null;column:group_name"`
	CountryLeft  string `json:"country_left" gorm:"type:varchar(64);not null;column:country_left"`   //比赛框左边的国家
	Score        string `json:"score" gorm:"type:varchar(64);not null;column:score"`                 //两边得分
	CountryRight string `json:"country_right" gorm:"type:varchar(64);not null;column:country_right"` //比赛框右边的国家
	MatchNumber  string `json:"match_number" gorm:"type:varchar(20);not null;column:match_number"`
	Referee      string `json:"referee" gorm:"referee"` //裁判
}

func (Match) TableName() string {
	return "matches"
}

type MatchSerializer struct {
	ID           uint       `json:"id"`
	CreatedAt    time.Time  `json:"create_at"`
	UpdateAt     time.Time  `json:"update_at"`
	DeleteAt     *time.Time `json:"delete_at"`
	Date         string     `json:"date"`
	GroupName    string     `json:"group_name"`
	Location     string     `json:"location"`
	CountryLeft  string     `json:"country_left"`
	CountryRight string     `json:"country_right"`
	Score        string     `json:"score"`
	MatchNumber  int        `json:"match_number"`
}
