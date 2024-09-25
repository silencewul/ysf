package ysf

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"time"
	"ysf/music"
)

type tResponse struct {
	Name             string  `json:"name"`
	SalesVol         float32 `json:"salesVol"`
	RemainStock      float32 `json:"remainStock"`
	SaleEnkStatus    int     `json:"saleEnkStatus"`
	Enable           int     `json:"enable"`
	SoldEnableStatus bool    `json:"soldEnableStatus"`
}

func TStart(num int) {
	local, _ := time.LoadLocation("Local")
	t := time.Now()
	// 判断是否加零
	tMonth, tDay := t.Month(), t.Day()
	var tMs, tDs = fmt.Sprintf("%d", tMonth), fmt.Sprintf("%d", tDay)
	if tMonth < 10 {
		tMs = fmt.Sprintf("0%d", tMonth)
	}
	if tDay < 10 {
		tDs = fmt.Sprintf("0%d", tDay)
	}
	timeStr := fmt.Sprintf("%d-%v-%v 08:10:00", t.Year(), tMs, tDs)
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, local) //2024-9-18 08:00:01
	//fmt.Println(startTime)

	for {

		// 判断时间
		for {
			if startTime.Before(time.Now()) {
				break
			}
			fmt.Println("还没到时间")
			time.Sleep(1000 * time.Millisecond)
		}

		c := colly.NewCollector(
			//Visit only domains hackerspaces.org, wiki.hackerspaces.org
			colly.UserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 16_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148/sa-sdk-ios  (com.unionpay.chsp) (cordova 4.5.4) (updebug 0) (version 1011) (UnionPay/1.0 CloudPay) (clientVersion 311) (language zh_CN) (languageFamily zh_CN) (upApplet single) (walletMode 00) "),
		)

		c.OnRequest(func(r *colly.Request) {
			//r.Headers.Set("cookie", "route=68507c350d203ede9aba382a4e62a92b;sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%2219127296f1b299-03e5fffd6347818-227f031e-329160-19127296f1c544%22%2C%22%24device_id%22%3A%2219127296f1b299-03e5fffd6347818-227f031e-329160-19127296f1c544%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC%22%2C%22%24latest_referrer_host%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfYW5vbnltb3VzX2lkIjoiMTkxMjcyOTZmMWIyOTktMDNlNWZmZmQ2MzQ3ODE4LTIyN2YwMzFlLTMyOTE2MC0xOTEyNzI5NmYxYzU0NCIsIiRpZGVudGl0eV9jb29raWVfaWQiOiIxOTE3ZGZlN2IxNDI3MzgtMGVjODc2OGRiYjA3ZWYtM2QyNTU3NDUtMzI5MTYwLTE5MTdkZmU3YjE1MmE1ZCJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%22%2C%22value%22%3A%22%22%7D%7D")
			r.Headers.Set("Accept", "*/*")
			r.Headers.Set("Accept-Encoding", "gzip, deflate, br")

			r.Headers.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")

			//r.Headers.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJ1aWQiOjE0MzMzMzc0MzkzNTUzOTMsInJvbGUiOiJVTklQQVlfVVNFUiIsInBydiI6IjUwMDIzZjllMDlmYWRiYzhjZDhhMmQ2NDBjZDkwZWI3OWUzYjQ2ZjgiLCJncmQiOiJvYXV0aCIsImlhdCI6MTcyNTc1NzI2MSwiZXhwIjoxNzI1ODQzNjYxfQ.BTlcU1pAGmlF1dKXvKXf9CLrc1x8K9x3PNHWxMgk-_I")

			r.Headers.Set("app", "cqylcs")
			r.Headers.Set("Referer", "https://cqbdshzq.cup.com.cn/wap/cqylcs/home/up/")
		})

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong2:", err)
		})

		c.OnResponse(func(r *colly.Response) {
			//判断code
			if r.StatusCode == 200 {
				body := new(tResponse)
				json.Unmarshal(r.Body, body)
				// 判断
				if body.RemainStock > 0 || body.Enable > 0 || body.SoldEnableStatus != false || body.SaleEnkStatus > 0 {
					log.Println(body)
					music.Play()
				}
				log.Println(body.Name, "--------------", body.SalesVol)
			}
		})

		//u10 := fmt.Sprintf("https://www.baidu.com")
		//u20 := fmt.Sprintf("https://www.baidu.com")

		u10 := fmt.Sprintf("https://cqbdshzq.cup.com.cn/wap/api-gateway/api-center/mall/goods/1452585911539617?relateId=-1&relateType=0")
		u20 := fmt.Sprintf("https://cqbdshzq.cup.com.cn/wap/api-gateway/api-center/mall/goods/1452586526463905?relateId=-1&relateType=0")

		url := []string{u10, u20}

		_ = c.Visit(url[num])

		time.Sleep(500 * time.Millisecond)

		//rand.Seed(time.Now().UnixNano())
		//var sleepTime = time.Duration(rand.Intn(4)) + 2
		//time.Sleep(sleepTime * time.Second)
	}

}
