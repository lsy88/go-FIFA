package domain

import (
	"FIFA/model"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

var (
	ErrorAwardUrl        = errors.New("get award url fail")
	ErrorAwardDownloader = errors.New("download award source fail")
)

func Awards(doc *goquery.Document) error {
	var err error
	//count := 0
	doc.Find(".fp-tournament-award-badge_awardContent__5BGw1").Each(func(i int, s *goquery.Selection) {
		s.Find(".fp-tournament-award-badge_awardName__a2jP2").Text()
		s.Find(".fp-tournament-award-badge_awardWinner__O1bxe").Text()
		s.Find(".fp-tournament-award-badge_awardWinnerCountry__octy_").Text()
	})
	
	// db save
	
	// push data into db
	//dao.DB.Save(&award)
	
	//fmt.Println(count)
	
	return err
}

func callBack(url string, doc *goquery.Document) []model.Award {
	
	allAwardInfo := make([]model.Award, 0, 0)
	
	awardName := doc.Find("h1").Eq(2).Text()
	
	doc.Find("div p").Each(func(i int, selection *goquery.Selection) {
		
		if i > 6 {
			
			awardInfo := selection.Text()
			if strings.HasPrefix(awardInfo, "*") {
				return
			}
			oneAward := model.Award{}
			oneAward.URL = url
			oneAward.AwardName = awardName
			oneAward.Info = awardInfo
			allAwardInfo = append(allAwardInfo, oneAward)
		}
		
	})
	return allAwardInfo
}
