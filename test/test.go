package test

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

type response struct {
	Data struct {
		Stock  int
		Status int
	}
	Msg  string
	Code int
}

func StartG() {

	var url = fmt.Sprintf("https://hd.netfishing.cn/index.php/shop/Goods/getGoodsDetail?goods_id=1")

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
				json.Unmarshal(r.Body, body)
				log.Println(body)
				if body.Data.Stock > 0 {
					wxPusher.Push("常州电影票")
					music.Play()
				}
				//fmt.Println(string(r.Body))
			}
		})

		_ = c.Visit(url)

		rand.Seed(time.Now().UnixNano())
		var sleepTime = time.Duration(rand.Intn(10)) + 10
		time.Sleep(sleepTime * time.Second)
	}

}
