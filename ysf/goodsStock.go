package ysf

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"time"
	"ysf/music"
	"ysf/wxPusher"
)
type responseN struct {
	Content struct{
		GoodsStockResDTOS []GoodsStockResDTOS
	}
	Status string
}

type GoodsStockResDTOS struct {
	GoodsInfoNo string
	Stock int
	SaleNum int

}
func StartG()  {

	//var htmlUrl = "https://mall.95516.com/newmobile/goods/202303013861001?channel=3"
//https://mall.95516.com/mobileapi/goods/goodsStock?goodsNo=202304014225&goodsInfoNo=202304014225001
	var goodsInfoNo = "202304014896001"

	var goodsNo = goodsInfoNo[:12]
	
	var url = fmt.Sprintf("https://mall.95516.com/mobileapi/goods/goodsStock?goodsNo=%s&goodsInfoNo=%s",goodsNo,goodsInfoNo)

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
				body := new(responseN)
				json.Unmarshal(r.Body,body)
				log.Println(body.Content.GoodsStockResDTOS[0].Stock)
				if body.Content.GoodsStockResDTOS[0].Stock > 0 {
					wxPusher.Push("渝中300-100")
					music.Play()
				}
				//fmt.Println(string(r.Body))
			}
		})

		_ = c.Visit(url)

		time.Sleep(500*time.Millisecond)
	}

}

