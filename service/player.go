package service

import (
	"FIFA/dao"
)

func GetPlayerById(id int) (player interface{}, err error) {
	return dao.GetPlayerById(id)
}

func GetPlayerList() (players interface{}, count int64, err error) {
	return dao.GetPlayerList()
}
