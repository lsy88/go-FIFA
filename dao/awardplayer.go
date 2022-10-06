package dao

import "FIFA/model"

func GetAwardPlayerList() (awardplayer []*model.AwardPlayer, count int64, err error) {
	err = DB.Model(&model.AwardPlayer{}).Count(&count).Find(&awardplayer).Error
	return
}

func GetAwardPlayerByKeyword(key string) (awardplayer []*model.AwardPlayer, count int64, err error) {
	err = DB.Model(&model.AwardPlayer{}).Where("award_name like ? or player_name like ? or country like ?", "%"+key+"%",
		"%"+key+"%", "%"+key+"%").Count(&count).Find(&awardplayer).Error
	return
}
