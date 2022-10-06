package service

import "FIFA/dao"

func GetAwardById(id int) (award interface{}, err error) {
	return dao.GetAwardById(id)
}

func GetAwardList() (awards interface{}, count int64, err error) {
	return dao.GetAwardList()
}

func GetAwardByKeyword(keyword string) (awards interface{}, count int64, err error) {
	return dao.GetAwardByKeyword(keyword)
}
