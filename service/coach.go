package service

import "FIFA/dao"

func GetCoachById(id int) (coach interface{}, err error) {
	return dao.GetCoachById(id)
}

func GetCoachList() (coaches interface{}, count int64, err error) {
	return dao.GetCoachList()
}
