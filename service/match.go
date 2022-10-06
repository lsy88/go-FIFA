package service

import "FIFA/dao"

func GetMatchList() (teams interface{}, count int64, err error) {
	return dao.GetMatchList()
}

func GetMatchById(id int) (team interface{}, err error) {
	return dao.GetMatchById(id)
}
