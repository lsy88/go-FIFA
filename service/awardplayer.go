package service

import "FIFA/dao"

func GetAwardPlayerList() (awardPlayer interface{}, count int64, err error) {
	return dao.GetAwardPlayerList()
}

func GetAwardPlayerByKeyword(key string) (awardPlayer interface{}, count int64, err error) {
	return dao.GetAwardPlayerByKeyword(key)
}
