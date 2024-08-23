package ysf

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"math/rand"
	"time"
	"ysf/music"
	"ysf/wxPusher"
)

type payload struct {
	Action       string   `json:"action"`
	UnionActCode []string `json:"union_act_code"`
}

type uresponse struct {
	Errmsg string `json:"errmsg"`
	Errno  string `json:"errno"`
	Data   []struct {
		DayRemainAmount        int    `json:"day_remain_amount"`
		DayRemainAmountPercent int    `json:"day_remain_amount_percent"`
		DayRemainCount         int    `json:"day_remain_count"`
		DayRemainCountPercent  int    `json:"day_remain_count_percent"`
		StartMark              int    `json:"start_mark"`
		UnionActCode           string `json:"union_act_code"`
	} `json:"data"`
}

func UStrat() {

	for {

		c := colly.NewCollector(
			//Visit only domains hackerspaces.org, wiki.hackerspaces.org
			//colly.UserAgent(Mozilla5.0 (Windows NT 6.1; WOW64) AppleWebKit537.36 (KHTML, like Gecko) Chrome81.0.4044.138 Safari537.36 NetTypeWIFI MicroMessenger7.0.20.1781(0x6700143B) WindowsWechat(0x63040026)),
			//colly.UserAgent(Mozilla5.0 (Windows NT 6.1) AppleWebKit537.36 (KHTML, like Gecko) Chrome41.0.2228.0 Safari537.36),
			colly.UserAgent("Mozilla5.0 (Windows NT 10.0; Win64; x64) AppleWebKit537.36 (KHTML, like Gecko) Chrome107.0.0.0 Safari537.36 MicroMessenger7.0.20.1781(0x6700143B) NetTypeWIFI MiniProgramEnvWindows WindowsWechatWMPF WindowsWechat(0x63090819)XWEB8461"),
		)

		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("Accept", "application/json, text/plain, */*")
			r.Headers.Set("Content-Type", "application/json;charset=UTF-8")

			r.Headers.Set("Sec-Fetch-Dest", "empty")
			r.Headers.Set("Sec-Fetch-Mode", "cors")
			r.Headers.Set("Sec-Fetch-Site", "cross-site")

			r.Headers.Set("xweb_xhr", "1")
			r.Headers.Set("Referer", "https://servicewechat.com/wx2610064b554eaff0/50/page-frame.html")
		})

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong2:", err)
		})

		c.OnResponse(func(r *colly.Response) {
			//判断code
			if r.StatusCode == 200 {
				body := new(uresponse)
				json.Unmarshal(r.Body, body)
				if body.Data[0].DayRemainCount > 0 || body.Data[0].DayRemainAmount > 0 {
					music.Play()
					wxPusher.Push("农行60-20")
				}

				log.Println(body.Data[0])
				fmt.Println(string(r.Body))
			}
		})

		url := fmt.Sprintf("https://umc.unionpay.com/smart/app/customerApi.php")

		data := payload{
			Action:       "app.home.HomeService.getCouponRemain",
			UnionActCode: []string{"3102024012561755"},
		}
		postData, _ := json.Marshal(data)
		_ = c.PostRaw(url, postData)

		rand.Seed(time.Now().UnixNano())
		var sleepTime = time.Duration(rand.Intn(4)) + 2
		time.Sleep(sleepTime * time.Second)
	}

}
