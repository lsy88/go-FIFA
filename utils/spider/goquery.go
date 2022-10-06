package spider

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strings"
	"time"
)

func DownloaderBySelenium(url string) (string, error) {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	imageCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imageCaps,
		Path:  "",
		Args: []string{
			"--headless", // debug使用,true不显示浏览器，false打开浏览器
			"--no-sandbox",
			"--User-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
			"--ignore-certificate-errors",
			"--ignore -ssl-errors",
		},
	}
	caps.AddChrome(chromeCaps)
	//启动Chromedriver
	service, err := selenium.NewChromeDriverService(
		"path/to/chromedriver", 8080, opts...,
	)
	defer service.Stop()

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	//调起Chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 8080))

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer webDriver.Quit()
	//导航到目标网站
	err = webDriver.Get(url)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	outputDiv, err := webDriver.FindElement(selenium.ByTagName, "span")

	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))

	//return webDriver.PageSource()
	return output, nil

}

// Downloader 使用传统的net/http爬取数据
func Downloader(url string) (*goquery.Document, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		zap.L().Error("http.NewRequest failed", zap.Error(err))
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return goquery.NewDocumentFromReader(response.Body)

}

//chromedp
func GetHttpHTML(url string, selector string) string {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false), //设置成无浏览器弹出
		chromedp.Flag("blink-settings", "imageEnable=false"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"),
	}
	c, _ := chromedp.NewExecAllocator(context.Background(), options...)
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	_ = chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	timeOutCtx, cancel := context.WithTimeout(chromeCtx, 90*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeOutCtx,
		chromedp.Navigate(url),
		//需要爬取的网页的url
		//chromedp.WaitVisible(`div[class="fp-tournament-award-badge_awardContent__dUtoO"]`),
		chromedp.WaitVisible(selector),
		//等待某个特定的元素出现
		chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
		//生成最终的html文件并保存在htmlContent文件中
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(htmlContent)
	return htmlContent
}

//在拿到的HTML中找到获奖球员的信息
func GetAwardPlayer(html string) ([][]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	//var award, name, country string
	awardTmp := [][]string{}
	doc.Find(`div[class="fp-tournament-award-badge_awardContent__dUtoO"]`).Each(func(i int, selection *goquery.Selection) {
		award := selection.Find(`p[class="fp-tournament-award-badge_awardName__JpsZZ"]`).Text()
		name := selection.Find(`h4[class=" fp-tournament-award-badge_awardWinner__P_z2d"]`).Text()
		country := selection.Find(`p[class="fp-tournament-award-badge_awardWinnerCountry__EmjVU"]`).Text()

		//fmt.Printf("%s--%s--%s\n", award, name, country)
		awardTmp = append(awardTmp, []string{award, name, country})
	})
	return awardTmp, err
}

func GetTeamMsg(html string) ([][]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	tmp := [][]string{}
	doc.Find(`div[class="fp-accordion-card_accordionCard__FpSE4"]`).Each(func(i int, selection *goquery.Selection) {
		var matchName, date, player, match, leftCountry, score, rightCountry string

		matchName = selection.Find(`p[class="ff-mb-0 fp-accordion-card_accordionTitle__RL0Vi"]`).Text()

		selection.Find(`div[class="fp-accordion-card_items__pqIbk"]`).Each(func(i int, body *goquery.Selection) {

			body.Find(`div[class="row fp-match-overview_matchRow__pYNtL justify-content-between d-flex"]`).Each(func(i int, line *goquery.Selection) {
				line.Find(`ul[class="fp-match-overview_matchInfoList__N1_cd"]`).Each(func(i int, datebody *goquery.Selection) {
					date = datebody.Find(`li[class="fp-match-overview_matchInfoListDate__dZ1Ms"]`).Text()
					player = datebody.Find(`li[class="fp-match-overview_matchInfoListDate__dZ1Ms"]+li`).Text()
					match = datebody.Find(`li[class="fp-match-overview_matchInfoListDate__dZ1Ms"]+li+li+li`).Text()
				})

				leftCountry = line.Find(`h4[class=" d-none d-md-inline col fp-match-overview_homeTeam__ceqhJ h4-article"]`).Text()

				line.Find(`div[class="ff-match-results_concludedMatchResults__tZXNL ff-match-results_results__357Cf ff-mx-md-24 ff-mx-lg-32 ff-mx-8 d-flex justify-content-center align-items-center"]`).Each(func(i int, center *goquery.Selection) {
					score = center.Find(`h4[class=" h4-article"]`).Text()
				})

				rightCountry = line.Find(`h4[class=" d-none d-md-inline col fp-match-overview_awayTeam__Dg9AY h4-article"]`).Text()
			})

		})
		tmp = append(tmp, []string{matchName, date, player, match, leftCountry, score, rightCountry})
	})
	return tmp, nil
}

func GetGroupMsg(html string) ([][]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	tmp := [][]string{}
	var groupName, num, cname, MP, W, D, L, GF, GA, jiajian, PTS string
	doc.Find(`div[class="fp-group-card_groupsTable__fGr6S"]>table>tbody`).Each(func(i int, selection *goquery.Selection) {
		groupName = selection.Find(`th[class="fp-group-card_alignLeft__XmFkr"]>div`).Text()

		selection.Find(`div[class="fp-group-card_groupsTable__fGr6S"]>table>tbody>tr+tr`).Each(func(i int, line *goquery.Selection) {

			num = line.Find(`div[class="fp-group-card_groupPosition__Vv_es"]`).Text()
			cname = line.Find(`td[class="fp-group-card_teamName__sNMrg fp-group-card_alignLeft__XmFkr"]`).Text()

			MP = line.Find(`td[class="fp-group-card_teamName__sNMrg fp-group-card_alignLeft__XmFkr"]+td`).Text()
			W = line.Find(`td[class="fp-group-card_teamName__sNMrg fp-group-card_alignLeft__XmFkr"]+td+td`).Text()
			D = line.Find(`td[class="fp-group-card_teamName__sNMrg fp-group-card_alignLeft__XmFkr"]+td+td+td`).Text()
			L = line.Find(`td[class="fp-group-card_teamName__sNMrg fp-group-card_alignLeft__XmFkr"]+td+td+td+td`).Text()
			GF = line.Find(`td[class="fp-group-card_teamName__sNMrg fp-group-card_alignLeft__XmFkr"]+td+td+td+td+td`).Text()
			GA = line.Find(`td[class="fp-group-card_teamName__sNMrg fp-group-card_alignLeft__XmFkr"]+td+td+td+td+td+td`).Text()
			jiajian = line.Find(`td[class="fp-group-card_teamName__sNMrg fp-group-card_alignLeft__XmFkr"]+td+td+td+td+td+td+td`).Text()
			PTS = line.Find(`td[class="fp-group-card_teamPoints__7apoM"]`).Text()
			tmp = append(tmp, []string{groupName, num, cname, MP, W, D, L, GF, GA, jiajian, PTS})
		})

	})

	return tmp, nil
}

//获取球员信息
func GetPlayer(html string) ([][]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	tmp := [][]string{}
	var xing, ming, seat, paiming, country string
	doc.Find(`div[class="fp-squad_squadCardsContainer__0PeCv"]`).Each(func(i int, selection *goquery.Selection) {

		selection.Find(`div[class="fp-squad-player-card_playerDetails__dMUgv"]`).Each(func(i int, player *goquery.Selection) {
			player.Find(`div[class="fp-squad-player-card_firstRow__4l1On"]`).Each(func(i int, name *goquery.Selection) {
				xing = name.Find(`div[class="fp-squad-player-card_firstName__o0cWG"]`).Text()
				ming = name.Find(`div[class="fp-squad-player-card_lastName__TNGsc"]`).Text()
				paiming = name.Find(`div[class="fp-squad-player-card_jerseyNumber__wfEsB"]`).Text()

			})

			player.Find(`div[class="fp-squad-player-card_secondRow__tWyB6"]`).Each(func(i int, s *goquery.Selection) {
				seat = s.Find(`div[class="fp-squad-player-card_position__UHe_f"]`).Text()
				s.Find(`div[class="fp-squad-player-card_flag__jMbx6"]`).Each(func(i int, img *goquery.Selection) {
					country, _ = img.Find(`img[class="image_img__jrck5"]`).Attr("alt")
				})
			})

			tmp = append(tmp, []string{xing + " " + string(ming[0]) + strings.ToLower(ming[1:]), seat, paiming, country})
		})

	})

	return tmp, nil
}

//获取队伍信息
func GetTeam(html string) ([][]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	tmp := [][]string{}
	var f, country string
	doc.Find(`div[class="row"]`).Each(func(i int, selection *goquery.Selection) {
		selection.Find(`div[class="ff-display-card_displayCard__xUHUW"]`).Each(func(i int, body *goquery.Selection) {
			body.Find(`div[class="ff-display-card_displayCardImage__ohjAR"]`).Each(func(i int, flag *goquery.Selection) {
				f, _ = flag.Find(`img[class="image_img__jrck5"]`).Attr("src")
				country, _ = flag.Find(`img[class="image_img__jrck5"]`).Attr("alt")

			})
			tmp = append(tmp, []string{f, country})
		})
	})
	return tmp, nil
}

//拿到比赛信息
func GetMatch(html string) ([][]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	tmp := [][]string{}
	var date, riqi, caipan, xiaozu, match, leftCountry, bifen, rightCountry string

	doc.Find(`div[class="fp-accordion-card_innerAccordion__fGvOm "]`).Each(func(i int, selection *goquery.Selection) {
		selection.Find(`div[class="fp-accordion-card_alignItems__cy1UV"]`).Each(func(i int, time *goquery.Selection) {
			date = time.Find(`p[class="ff-mb-0 fp-accordion-card_accordionTitle__RL0Vi"]`).Text()
		})

		selection.Find(`div[class="fp-accordion-card_items__pqIbk"]`).Each(func(i int, relay *goquery.Selection) {
			relay.Find(`div[class="row fp-match-overview_matchRow__pYNtL justify-content-between d-flex"]`).Each(func(i int, detail *goquery.Selection) {
				detail.Find(`div[class="d-none d-md-inline col fp-match-overview_tournamentDetails__we3GP"]`).Each(func(i int, t *goquery.Selection) {
					riqi = t.Find(`li[class="fp-match-overview_matchInfoListDate__dZ1Ms"]`).Text()
					caipan = t.Find(`li[class="fp-match-overview_matchInfoListDate__dZ1Ms"]+li`).Text()
					xiaozu = t.Find(`li[class="fp-match-overview_matchInfoListDate__dZ1Ms"]+li+li`).Text()
					match = t.Find(`li[class="fp-match-overview_matchInfoListDate__dZ1Ms"]+li+li+li`).Text()
				})

				leftCountry = detail.Find(`h4[class=" d-none d-md-inline col fp-match-overview_homeTeam__ceqhJ h4-article"]`).Text()

				detail.Find(`div[class="col-auto fp-match-overview_centerItems__gkRTw ff-mx-md-24 ff-mx-lg-32 ff-mx-sm-8"]`).Each(func(i int, w *goquery.Selection) {
					w.Find(`div[class="ff-match-results_concludedMatchResults__tZXNL ff-match-results_results__357Cf ff-mx-md-24 ff-mx-lg-32 ff-mx-8 d-flex justify-content-center align-items-center"]`).Each(func(i int, n *goquery.Selection) {
						bifen = n.Find(`h4[class=" h4-article"]`).Text()
					})
				})

				rightCountry = detail.Find(`h4[class=" d-none d-md-inline col fp-match-overview_awayTeam__Dg9AY h4-article"]`).Text()
				tmp = append(tmp, []string{date, riqi, caipan, xiaozu, match, leftCountry, bifen, rightCountry})
			})
		})
	})
	return tmp, nil
}
