package main

import (
	"FIFA/core"
	"FIFA/dao"
	"FIFA/router"
	"fmt"
)

// @title FIFA
// @version 1.0
// @description 世界杯后台系统
// @termsOfService http://swagger.io/terms/
// @host http://127.0.0.1:8080/FIFA/v1/api
func main() {
	if err := core.ViperInit(); err != nil {
		fmt.Printf("load config viper failed, err:#{err}\n")
		return
	}
	if err := core.ZapInit(core.Conf.LogConfig, core.Conf.Mode); err != nil {
		fmt.Printf("load config zap failed, err:#{err}\n")
		return
	}
	
	if err := dao.Init(); err != nil {
		return
	}
	
	//TODO 爬取数据入库
	//html, err := spider.GetHttpHTML(core.Conf.URL,
	//	`#content > div > div.fp-tournament-matches-overview_matchesContainer__iaXdn > div > div > div.d-none.d-xxl-block > div > div > div.fp-brackets-view_tournamentBrackets__icdaw.justify-content-center > ul.fp-brackets-view_left__h2_qA.fp-brackets-view_bracket__A_MNw.fp-brackets-view_roundOf16__e4Iyr > li:nth-child(1) > a > div > div.fp-match-card_matchInfo__Kjs5W`)
	//if err != nil {
	//	zap.L().Error("", zap.Error(err))
	//	return
	//}
	
	//html := spider.GetHttpHTML("https://www.fifa.com/tournaments/mens/worldcup/2018russia/teams/43935",
	//	`#content > div > div.fp-squad_squadContainer__PpIp8 > div > div.fp-squad_squadCardsContainer__0PeCv`)
	//
	////if err != nil {
	////	zap.L().Error("", zap.Error(err))
	////	return
	////}
	//fmt.Println(html)
	//award, err := spider.GetAwardPlayer(html)
	//if err != nil {
	//return
	//}
	//fmt.Println(award)
	
	//fmt.Println(html)
	//match, err := spider.GetPlayer(html)
	//if err != nil {
	//	return
	//}
	//fmt.Println("-------------------------------------")
	//fmt.Println(match)
	//爬取球员信息入库
	//for _, v := range match {
	//	m := model.Player{
	//		Name:    v[0],
	//		Role:    v[1],
	//		Number:  v[2],
	//		Country: v[3],
	//	}
	//	dao.DB.Model(&model.Player{}).Create(&m)
	//}
	//fmt.Println(match)
	//爬取比赛信息入库
	//for _, v := range match {
	//	m := model.Match{
	//		GroupName:    v[0],
	//		Date:         v[1],
	//		Referee:      v[2],
	//		MatchNumber:  v[3],
	//		CountryLeft:  v[4],
	//		Score:        v[5],
	//		CountryRight: v[6],
	//	}
	//	dao.DB.Model(&model.Match{}).Create(&m)
	//}
	
	//html := spider.GetHttpHTML("https://www.fifa.com/tournaments/mens/worldcup/2018russia/teams",
	//	`#content > div > div:nth-child(3) > div > div:nth-child(2)`)
	//team, _ := spider.GetTeam(html)
	//fmt.Println(team)
	//for _, v := range team {
	//	m := model.Team{
	//		Flag: v[0],
	//		Name: v[1],
	//	}
	//	dao.DB.Model(&model.Team{}).Create(&m)
	//}
	
	//拿到比赛信息入库更新版
	//html := spider.GetHttpHTML("https://www.fifa.com/tournaments/mens/worldcup/2018russia/match-center",
	//	`#content > div > div.ff-bg-grey-lightest > div.fp-tournament-matches-overview_matchesContainer__iaXdn > div > div > div:nth-child(2) > div > div`)
	//group, _ := spider.GetMatch(html)
	////fmt.Println(group)
	//for _, v := range group {
	//	m := model.Match{
	//		Date:         v[1],
	//		GroupName:    v[0],
	//		Referee:      v[2],
	//		MatchNumber:  v[4],
	//		CountryLeft:  v[5],
	//		Score:        v[6],
	//		CountryRight: v[7],
	//	}
	//	dao.DB.Model(&model.Match{}).Create(&m)
	//}
	
	//爬取小组赛信息入库
	//html := spider.GetHttpHTML("https://www.fifa.com/tournaments/mens/worldcup/2018russia/match-center",
	//	`#content > div > div.ff-bg-grey-lightest > div.ff-pt-32.ff-pt-lg-48 > div > div.fp-groups-overview_carouselInner__YHac6 > div > div > div > div`)
	//group, _ := spider.GetGroupMsg(html)
	//fmt.Println(group)
	//for _, v := range group {
	//	if v[2] == "" {
	//		continue
	//	}
	//	v3, _ := strconv.Atoi(v[3])
	//	v4, _ := strconv.Atoi(v[4])
	//	v5, _ := strconv.Atoi(v[5])
	//	v6, _ := strconv.Atoi(v[6])
	//	v7, _ := strconv.Atoi(v[7])
	//	v8, _ := strconv.Atoi(v[8])
	//	v9, _ := strconv.Atoi(v[9])
	//	v10, _ := strconv.Atoi(v[10])
	//
	//	m := model.Group{
	//		GroupName:   v[0],
	//		TeamName:    v[2][1:],
	//		MatchPlayed: v3,
	//		WinNumber:   v4,
	//		Draw:        v5,
	//		Lost:        v6,
	//		GoalFor:     v7,
	//		GoalAgainst: v8,
	//		DiffGoal:    v9,
	//		Points:      v10,
	//	}
	//	dao.DB.Model(&model.Group{}).Create(&m)
	//}
	
	router.SetupRouter(core.Conf.Mode)
}
