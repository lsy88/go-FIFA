package model

import (
	"gorm.io/gorm"
)

//获得金牌的团队统计
type TeamStatisticWithTopGoal struct {
	gorm.Model
	Rank          int    `json:"rank" gorm:"type:integer; not null; column:rank"`
	TeamName      string `json:"team_name" gorm:"type:varchar(12); not null; column:team_name"`
	MatchPlayed   int    `json:"match_played" gorm:"type:integer; not null; column:matches_payed"`
	GoalsFor      int    `json:"goals_for" gorm:"type:integer; not null; column:goals_for"`
	GoalsScored   int    `json:"goals_scored" gorm:"type:integer; not null; column:goals_scored"`
	GoalsAgainst  int    `json:"goals_against" gorm:"type:integer; not null; column:goals_against"`
	PenaltyGoal   int    `json:"penalty_goal" gorm:"type:integer; not null; column:penalty_goal"`
	OwnGoals      int    `json:"own_goals" gorm:"type:integer; not null; column:own_goals"`
	OpenPlayGoals int    `json:"open_play_goals" gorm:"type:integer; not null; column:open_play_goals"`
	SetPieceGoals int    `json:"set_piece_goals" gorm:"type:integer; not null; column:set_piece_goals"`
}

//队伍参赛次数统计
type TeamStatisticWithAttempts struct {
	gorm.Model
	Rank              int    `json:"rank" gorm:"type:integer; not null; column:rank"`
	TeamName          string `json:"team_name" gorm:"type:varchar(12); not null; column:team_name"`
	MatchPlayed       int    `json:"match_played" gorm:"type:integer; not null; column:matches_payed"`
	Shots             int    `json:"shots" gorm:"type:integer; not null; column:shots_number"`
	AttemptsOnTarget  int    `json:"attempts_on_target" gorm:"type:integer; not null; column:attempts_on_target"`
	AttemptsOffTarget int    `json:"attempts_off_target" gorm:"type:integer; not null; column:attempts_off_target"`
	ShotsBlocked      int    `json:"shots_blocked" gorm:"type:integer; not null; column:shots_blocked_number"`
}

//队伍违纪次数
type TeamStatisticWithDisciplinary struct {
	gorm.Model
	Rank                int    `json:"rank" gorm:"type:integer; not null; column:rank"`
	TeamName            string `json:"team_name" gorm:"type:varchar(12); not null; column:team_name"`
	MatchPlayed         int    `json:"match_played" gorm:"type:integer; not null; column:matches_payed"`
	YellowCards         int    `json:"yellow_cards" gorm:"type:integer; default(0); column:yellow_cards"`
	SecondYellow        int    `json:"second_yellow" gorm:"type:integer; default(0); column:second_yellow_cards"`
	RedCards            int    `json:"red_cards" gorm:"type:integer; default(0); column:red_cards"`
	FoulsCommitted      int    `json:"fouls_committed" gorm:"type:integer; default(0); column:fouls_committed"`
	FoulsSuffered       int    `json:"fouls_suffered" gorm:"type:integer; default(0); column:fouls_suffered"`
	FoulsCausingPenalty int    `json:"fouls_causing_penalty" gorm:"type:integer; default(0); column:fouls_causing_penalty"`
}

//运动员进球信息统计
type PlayersStatisticWithGoalsScored struct {
	gorm.Model
	Rank             int    `json:"rank" gorm:"type:integer; not null; column:rank"`
	PlayerName       string `json:"player_name" gorm:"type:varchar(32); not null; column:player_name"`
	GoalsScored      int    `json:"goals_scored" gorm:"type:integer; not null; column:goals_scored"` //进球得分
	Assists          int    `json:"assists" gorm:"type:integer; not null; column:assists"`           //助攻得分
	MinutesPlayed    int    `json:"minutes_played" gorm:"type:integer; not null; column:minutes_played"`
	MatchesPlayed    int    `json:"matches_played" gorm:"type:integer; not null; column:matches_played"`          //出场时间
	PenaltiesScored  int    `json:"penalties_scored" gorm:"type:integer; not null; column:penalties_scored"`      //比赛场次
	GoalsScoredLeft  int    `json:"goals_scored_left" gorm:"type:integer; default:0; column:goals_scored_left"`   //左路进球数
	GoalsScoredRight int    `json:"goals_scored_right" gorm:"type:integer; default:0; column:goals_scored_right"` //右路进球数
	HeadedGoals      int    `json:"headed_goals" gorm:"type:integer; default:0; column:headed_goals"`             //头球
}

//最佳救球
type PlayersStatisticWithTopSave struct {
	gorm.Model
	Rank          int    `json:"rank" gorm:"type:integer; not null; column:rank"` //排名
	PlayerName    string `json:"player_name" gorm:"type:varchar(32); not null; column:player_name"`
	MatchedPlayed int    `json:"matched_played" gorm:"type:integer; column:matched_played"` //比赛场次
	MinutesPlayed int    `json:"minutes_played" gorm:"type:integer; column:minutes_played"` //出场时间
	Saves         int    `json:"saves" gorm:"type:integer; default(0); column:saves_number"`
	SaveRate      string `json:"save_rate" gorm:"type:varchar(20); column:save_rate"`
}

//获得最佳射手的球员
type PlayersStatisticWithShot struct {
	gorm.Model
	Rank              int    `json:"rank" gorm:"type:integer; not null; column:rank"`
	Player            string `json:"player" gorm:"type:varchar(32); not null; column:player_name"`
	MatchesPlayed     int    `json:"matches_played" gorm:"type:integer;column:matches_played"`
	MinutesPlayed     int    `json:"minutes_played" gorm:"type:integer;column:minutes_played"`
	Shots             int    `json:"shots" gorm:"type:integer;column:shots_number"`
	AttemptsOnTarget  int    `json:"attempts_on_target" gorm:"type:integer;column:attempts_On_target"`   //射中次数
	AttemptsOffTarget int    `json:"attempts_off_target" gorm:"type:integer;column:attempts_off_target"` //不中次数
	ShotsBlocked      int    `json:"shots_blocked" gorm:"type:integer;column:shots_blocked"`             //射门被挡次数
}

//球员违规次数
type PlayersStatisticWithDisciplinary struct {
	gorm.Model
	Rank                int    `json:"rank" gorm:"type:integer; not null; column:rank"`
	PlayerName          string `json:"player_name" gorm:"type:varchar(64); not null; column:player_name"`
	MatchPlayed         int    `json:"match_played" gorm:"type:integer;column:matches_payed"`
	YellowCards         int    `json:"yellow_cards" gorm:"type:integer;column:yellow_cards"`
	SecondYellow        int    `json:"second_yellow" gorm:"type:integer;column:second_yellow_cards"`
	RedCards            int    `json:"red_cards" gorm:"type:integer;column:red_cards"`
	FoulsCommitted      int    `json:"fouls_committed" gorm:"type:integer;column:fouls_committed"`
	FoulsSuffered       int    `json:"fouls_suffered" gorm:"type:integer;column:fouls_suffered"`
	FoulsCausingPenalty int    `json:"fouls_causing_penalty" gorm:"type:integer;column:fouls_causing_penalty"`
}
