package service

import (
	"FIFA/controller/request"
	"FIFA/dao"
	"FIFA/model"
)

func GetPlayerGoalByRank(rank int) (player interface{}, err error) {
	return dao.GetPlayerGoalByRank(rank)
}

func GetPlayerGoalHandler(req *request.ReqPlayerGoal) (player []*model.PlayersStatisticWithGoalsScored, count int64, err error) {
	if req.Name != "" {
		return dao.GetPlayerGoalName(req.Name)
	}
	if req.Goal != "" {
		return dao.GetPlayerGoalScore(req.Goal)
	}
	return dao.GetPlayerGoalAll()
}

func GetPlayerSaveByRank(rank int) (player interface{}, err error) {
	return dao.GetPlayerSaveByRank(rank)
}

func GetPlayerSave(name string) (player interface{}, count int64, err error) {
	return dao.GetPlayerSave(name)
}

func GetPlayerShotByRank(rank int) (player interface{}, err error) {
	return dao.GetPlayerShotByRank(rank)
}

func GetPlayerShot(name string) (player interface{}, count int64, err error) {
	return dao.GetPlayerShot(name)
}

func GetPlayerDisciplinaryByRank(rank int) (player interface{}, err error) {
	return dao.GetPlayerDisciplinaryByRank(rank)
}
func GetPlayerDisciplinary() (player interface{}, count int64, err error) {
	return dao.GetPlayerDisciplinary()
}

func GetTeamGoalByRank(rank int) (team interface{}, err error) {
	return dao.GetTeamGoalByRank(rank)
}

func GetTeamGoal() (team interface{}, count int64, err error) {
	return dao.GetTeamGoal()
}

func GetTeamAttemptsByRank(rank int) (team interface{}, err error) {
	return dao.GetTeamAttemptsByRank(rank)
}

func GetTeamAttempts() (team interface{}, count int64, err error) {
	return dao.GetTeamAttempts()
}

func GetTeamDisciplinaryByRank(rank int) (team interface{}, err error) {
	return dao.GetTeamDisciplinaryByRank(rank)
}

func GetTeamDisciplinary() (team interface{}, count int64, err error) {
	return dao.GetTeamDisciplinary()
}
