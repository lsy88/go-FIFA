package service

import (
	"FIFA/controller/request"
	"FIFA/dao"
)

func GetClassicByYear(year int) (classic interface{}, err error) {
	return dao.GetClassicByYear(year)
}

func GetClassicList(req *request.ReqClassicMes) (classics interface{}, count int64, err error) {
	if req.AllList {
		classics, count, err = dao.GetClassicAll()
		return
	}
	if req.Title != "" {
		classics, count, err = dao.GetClassicByTitle(req.Title)
		return
	}
	if req.FinalResult != "" {
		classics, count, err = dao.GetClassicByResult(req.FinalResult)
		return
	}
	return
}
