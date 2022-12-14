package model

import (
	"gorm.io/gorm"
	"time"
)

//小组赛
type Group struct {
	gorm.Model
	GroupName   string `gorm:"type:varchar(64);not null;column:group_name"`
	TeamName    string `gorm:"type:varchar(64);not null;column:team_name"` //参赛队伍
	MatchPlayed int    `gorm:"type:integer;not null;column:match_played"`
	WinNumber   int    `gorm:"type:integer;not null;column:win_number"` //得分
	Draw        int    `gorm:"type:integer;not null;column:draw_number"`
	Lost        int    `gorm:"type:integer;not null;column:lost_number"`
	GoalFor     int    `gorm:"type:integer;not null;column:goal_number"`
	GoalAgainst int    `gorm:"type:integer;not null;column:goal_against"`
	DiffGoal    int    `gorm:"type:integer;not null;column:diff_goal"`
	Points      int    `gorm:"type:integer;not null;column:points"`
}

type GroupSerializer struct {
	ID          uint       `json:"id"`
	CreateAt    time.Time  `json:"create_at"`
	UpdateAt    time.Time  `json:"update_at"`
	DeleteAt    *time.Time `json:"delete_at"`
	GroupName   string     `json:"group_name"`
	TeamName    string     `json:"team_name"`
	MatchPlayed int        `json:"match_played"`
	WinNumber   int        `json:"win_number"`
	Draw        int        `json:"draw_number"`
	Lost        int        `json:"lost_number"`
	GoalFor     int        `json:"goal_number"`
	GoalAgainst int        `json:"goal_against"`
	DiffGoal    int        `json:"diff_goal"`
	Points      int        `json:"points"`
}

func (g *Group) Serializer() GroupSerializer {
	return GroupSerializer{
		ID:          g.ID,
		CreateAt:    g.CreatedAt,
		UpdateAt:    g.UpdatedAt,
		DeleteAt:    &g.DeletedAt.Time,
		GroupName:   g.GroupName,
		TeamName:    g.TeamName,
		MatchPlayed: g.MatchPlayed,
		WinNumber:   g.WinNumber,
		Draw:        g.Draw,
		Lost:        g.Lost,
		GoalFor:     g.GoalFor,
		GoalAgainst: g.GoalAgainst,
		DiffGoal:    g.DiffGoal,
		Points:      g.Points,
	}
}