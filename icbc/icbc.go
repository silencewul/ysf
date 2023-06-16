package icbc

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"time"
	"ysf/music"
)
type response struct {
	Data []ResponseData
	Message string
	Status int
}

type ResponseData struct {
	Gid int
	Num int
}
func Start()  {

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
				body := new(response)
				json.Unmarshal(r.Body,body)
				log.Println(body)
				if body.Data[7].Num > 0 {
					music.Play()
				}
				//fmt.Println(string(r.Body))
			}
		})
		//https://content.95516.com/UCIWeb/outputWeb/coupon/coupon.html#/couponDetail?couponId=3102023031363076
		url := fmt.Sprintf("https://gdecard.jiahuaming.com/apiv4/goods/getStocks?aid=906&city=%s","%E4%B8%8A%E6%B5%B7")
		_ = c.Visit(url)

		time.Sleep(1*time.Second)
	}

}

