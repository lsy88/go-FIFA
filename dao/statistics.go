package dao

import (
	"FIFA/model"
	"strconv"
)

func GetPlayerGoalByRank(rank int) (player *model.PlayersStatisticWithGoalsScored, err error) {
	err = DB.Model(&model.PlayersStatisticWithGoalsScored{}).Where("rank = ?", rank).First(&player).Error
	return
}

//根据姓名查询
func GetPlayerGoalName(name string) (player []*model.PlayersStatisticWithGoalsScored, count int64, err error) {
	err = DB.Model(&model.PlayersStatisticWithGoalsScored{}).Where("player_name = ?", name).Count(&count).Find(&player).Error
	return
}

//根据得分查询
func GetPlayerGoalScore(goal string) (player []*model.PlayersStatisticWithGoalsScored, count int64, err error) {
	score, _ := strconv.Atoi(goal)
	err = DB.Model(&model.PlayersStatisticWithGoalsScored{}).Where("goals_scored = ?", score).Count(&count).Find(&player).Error
	return
}

//查询出所有数据
func GetPlayerGoalAll() (player []*model.PlayersStatisticWithGoalsScored, count int64, err error) {
	err = DB.Model(&model.PlayersStatisticWithGoalsScored{}).Count(&count).Find(&player).Error
	return
}

//根据最佳救球排名查询
func GetPlayerSaveByRank(rank int) (player *model.PlayersStatisticWithTopSave, err error) {
	err = DB.Model(&model.PlayersStatisticWithTopSave{}).Where("rank = ?", rank).First(&player).Error
	return
}

//根据姓名查询,如果没有姓名就返回全部
func GetPlayerSave(name string) (player []*model.PlayersStatisticWithTopSave, count int64, err error) {
	if name == "" {
		err = DB.Model(&model.PlayersStatisticWithTopSave{}).Count(&count).Find(&player).Error
	} else {
		err = DB.Model(&model.PlayersStatisticWithTopSave{}).Where("name = ?", name).Count(&count).Find(&player).Error
	}
	return
}

func GetPlayerShotByRank(rank int) (player *model.PlayersStatisticWithShot, err error) {
	err = DB.Model(&model.PlayersStatisticWithShot{}).Where("rank = ?", rank).First(&player).Error
	return
}

//根据姓名查询射门情况
func GetPlayerShot(name string) (player []*model.PlayersStatisticWithShot, count int64, err error) {
	if name == "" {
		err = DB.Model(&model.PlayersStatisticWithShot{}).Count(&count).Find(&player).Error
		
	} else {
		err = DB.Model(&model.PlayersStatisticWithShot{}).Where("name = ?", name).Count(&count).Find(&player).Error
	}
	return
}

//根据排名查看球员违规情况
func GetPlayerDisciplinaryByRank(rank int) (player *model.PlayersStatisticWithDisciplinary, err error) {
	err = DB.Model(&model.PlayersStatisticWithDisciplinary{}).Where("rank = ?", rank).First(&player).Error
	return
}

func GetPlayerDisciplinary() (player []*model.PlayersStatisticWithDisciplinary, count int64, err error) {
	err = DB.Model(&model.PlayersStatisticWithDisciplinary{}).Count(&count).Find(&player).Error
	return
}

func GetTeamGoalByRank(rank int) (team *model.TeamStatisticWithTopGoal, err error) {
	err = DB.Model(&model.TeamStatisticWithTopGoal{}).Where("rank = ?", rank).First(&team).Error
	return
}

func GetTeamGoal() (team []*model.TeamStatisticWithTopGoal, count int64, err error) {
	err = DB.Model(&model.TeamStatisticWithTopGoal{}).Count(&count).Find(&team).Error
	return
}

func GetTeamAttemptsByRank(rank int) (team *model.TeamStatisticWithAttempts, err error) {
	err = DB.Model(&model.TeamStatisticWithAttempts{}).Where("rank = ?", rank).First(&team).Error
	return
}

func GetTeamAttempts() (team []*model.TeamStatisticWithAttempts, count int64, err error) {
	err = DB.Model(&model.TeamStatisticWithAttempts{}).Count(&count).Find(&team).Error
	return
}

func GetTeamDisciplinaryByRank(rank int) (team *model.TeamStatisticWithDisciplinary, err error) {
	err = DB.Model(&model.TeamStatisticWithDisciplinary{}).Where("rank = ?", rank).First(&team).Error
	return
}

func GetTeamDisciplinary() (team []*model.TeamStatisticWithDisciplinary, count int64, err error) {
	err = DB.Model(&model.TeamStatisticWithDisciplinary{}).Count(&count).Find(&team).Error
	return
}
