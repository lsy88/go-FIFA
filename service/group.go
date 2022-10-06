package service

import "FIFA/dao"

func GetGroupById(id int) (group interface{}, err error) {
	return dao.GetGroupById(id)
}

func GetGroupList() (groups interface{}, count int64, err error) {
	return dao.GetGroupList()
}
