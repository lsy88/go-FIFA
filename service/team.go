package service

import "FIFA/dao"

func GetTeamList() (teams interface{}, count int64, err error) {
	return dao.GetTeamList()
}

func GetTeamById(id int) (team interface{}, err error) {
	return dao.GetTeamById(id)
}
