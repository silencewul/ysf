package ysf

import (
"encoding/json"
"fmt"
"github.com/gocolly/colly"
"log"
"time"
	"ysf/music"
)
type responseT struct {
	Data struct{
		DialogVo DialogVo
	}
	Status string
	Message string
}

type DialogVo struct {
	GoodsInfoStock int
}
func StartT()  {

	for {
		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
			colly.UserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"),
		)

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong:", err)
		})

		c.OnResponse(func(r *colly.Response) {
			// 判断code
			if r.StatusCode == 200 {
				body := new(responseT)
				json.Unmarshal(r.Body,body)
				log.Println(body)
				if body.Data.DialogVo.GoodsInfoStock > 0 {
					music.Play()
				}
				//fmt.Println(string(r.Body))
			}
		})
//https://content.95516.com/UCIWeb/outputWeb/coupon/coupon.html#/couponDetail?couponId=3102023031363076
		url := fmt.Sprintf("https://mall.95516.com/mobile/api/productdetail.html?productId=10273151")
		_ = c.Visit(url)

		time.Sleep(5*time.Second)
	}

}

