package nsh

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"
	"time"
	"ysf/music"
	"ysf/wxPusher"
)
type response struct {
	Data struct{
		HasGrant bool `json:"has_grant"`
		List []PrizeList `json:"list"`
		ScoreCost string `json:"score_cost"`
	}
	ErrCode int
	Message string
	Timestamp int
}

type PrizeList struct {
	AwardName string `json:"award_name"`
	IsGrand string `json:"is_grand"`
	PrizeId  string `json:"prize_id"`
	PrizeName string `json:"prize_name"`
	TotalQuantity string `json:"total_quantity"`
}
func Start()  {

	//local ,_ := time.LoadLocation("Local")
	//endTime,_ := time.ParseInLocation("2006-01-02 15:04:05","2022-03-25 19:00:00",local)

	current := 0  //当前数量
	isPush := false  //是否已经推送消息

	randString := "d1f6cd47b320677be2446d33436c1393"

	for {
		//if endTime.Before(time.Now())  {
		//	break
		//}
			c := colly.NewCollector(
				// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
				colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63040026)"),
				//colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
			)

			c.OnRequest(func(r *colly.Request) {
				r.Headers.Set("cookie","_LOGIN_WX_UID_wx7dfb76a2f5f50323=d0VuRVJRam5rcVB2VHhJdGtwS1VyZz09")
				r.Headers.Set("Referer",fmt.Sprintf("http://fhsw.%s.event.linktovip.com/",randString))
			})

			c.OnError(func(_ *colly.Response, err error) {
				log.Println("Something went wrong:", err)
			})

			c.OnResponse(func(r *colly.Response) {
				// 判断code
				if r.StatusCode == 200 {
					body := new(response)
					json.Unmarshal(r.Body,body)
					log.Println(body)
					// 判断是否有88红包
					for _,v := range body.Data.List {
						prizeId,_ := strconv.Atoi(v.PrizeId)
						totalQuantity,_ := strconv.Atoi(v.TotalQuantity)
						// 判断是否发送消息
						if prizeId == 45 && totalQuantity > 10000 && totalQuantity != current {
							if isPush == false {
								wxPusher.Push(v.PrizeName)
								isPush = true
							}
							current = totalQuantity
							music.Play()
						}
						if prizeId == 45 && totalQuantity < 10000  {
							isPush = false
						}

						// 检测手表
						if prizeId == 56 && totalQuantity > 10000 {
							wxPusher.Push(v.PrizeName)
							music.Play()
						}
					}

					//fmt.Println(string(r.Body))
				}
			})

			url := fmt.Sprintf("http://fhsw.%s.event.linktovip.com/event/lottery/prizes",randString)
			_ = c.Visit(url)

			time.Sleep(10*time.Second)
	}

}
