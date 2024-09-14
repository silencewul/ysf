package ysf

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"time"
	"ysf/wxPusher"
)

type coupon struct {
	Title    string
	CouponId string
}

type response struct {
	Resp   string
	Msg    string
	Params struct {
		CouponId           string
		CouponQuota        string
		CouponQuotaLow     string
		HasRemainQuota     string
		QuotaStatus        string
		CouponQuotaPercent string
		CouponQuotaUseOut  string
	}
}

var couponList = make([]coupon, 0)

func Start() {
	couponList = []coupon{
		//{Title: "优惠日山姆刷卡满1019元减109元优惠", CouponId: "3102023110736199"},
		//{Title: "永悦消费乐享无限-满300元减100元消费券", CouponId: "3102023092421932"},
		//{Title: "约惠两江社保卡满300元减100元通用消费券", CouponId: "3102024021967275"},
		//{Title: "农行借记卡中石油刷卡满200元减20元活动", CouponId: "3102023102531589"},
		{Title: "高新区金秋消费季一期零售类满1000减300元", CouponId: "3102024091227040"},
		{Title: "高新区金秋消费季一期零售类满500减150元", CouponId: "3102024091227039"},
	}

	fmt.Println("请选择你要监控的卡券")

	for i, v := range couponList {
		fmt.Println(i+1, v.Title)
	}
	var num int
	var count int
	_, err := fmt.Scanf("%d", &num)
	if err != nil || num < 1 || num > len(couponList) {
		log.Fatal("请输入1-", len(couponList), "数字")
	}
	var couponDetail = couponList[num-1]
	var sleepTime = 1000 * time.Millisecond

	//go func() {
	//	couponDetail := couponList[0]
	//	for {
	//		c := colly.NewCollector(
	//			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//			colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	//		)
	//
	//		c.OnError(func(_ *colly.Response, err error) {
	//			log.Println("Something went wrong1:", err)
	//		})
	//
	//		c.OnResponse(func(r *colly.Response) {
	//			// 判断code
	//			if r.StatusCode == 200 {
	//				body := new(response)
	//				json.Unmarshal(r.Body, body)
	//				if body.Params.HasRemainQuota != "0" {
	//					//wxPusher.Push(couponDetail.Title)
	//					play("./2.mp3")
	//					log.Println(couponDetail.Title, "-----", body.Params.CouponQuota)
	//				} else {
	//					//log.Println(couponDetail.Title, "-----", body.Params.CouponQuota)
	//				}
	//			}
	//		})
	//		//https://content.95516.com/UCIWeb/outputWeb/coupon/coupon.html#/couponDetail?couponId=3102023031363076
	//		url := fmt.Sprintf("https://content.95516.com/koala-pre/koala/coupon/state?couponId=%s&cityCd=500000", couponDetail.CouponId)
	//		_ = c.Visit(url)
	//		time.Sleep(sleepTime)
	//		//fmt.Println("---------------------------------")
	//	}
	//}()

	//go func() {
	//	couponDetail := couponList[1]
	//	for {
	//		c := colly.NewCollector(
	//			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//			colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	//		)
	//
	//		c.OnError(func(_ *colly.Response, err error) {
	//			log.Println("Something went wrong:", err)
	//		})
	//
	//		c.OnResponse(func(r *colly.Response) {
	//			// 判断code
	//			if r.StatusCode == 200 {
	//				body := new(response)
	//				json.Unmarshal(r.Body, body)
	//				if body.Params.HasRemainQuota != "0" {
	//					//wxPusher.Push(couponDetail.Title)
	//					play("./2.mp3")
	//					log.Println(couponDetail.Title, "-----", body.Params.CouponQuota)
	//				} else {
	//					//log.Println(couponDetail.Title, "-----", body.Params.CouponQuota)
	//				}
	//			}
	//		})
	//		//https://content.95516.com/UCIWeb/outputWeb/coupon/coupon.html#/couponDetail?couponId=3102023031363076
	//		url := fmt.Sprintf("https://content.95516.com/koala-pre/koala/coupon/state?couponId=%s&cityCd=500000", couponDetail.CouponId)
	//		_ = c.Visit(url)
	//		time.Sleep(sleepTime)
	//		//fmt.Println("---------------------------------")
	//	}
	//}()

	for {
		count++
		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
			colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		)

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong2:", err)
		})

		c.OnResponse(func(r *colly.Response) {
			// 判断code
			if r.StatusCode == 200 {
				body := new(response)
				json.Unmarshal(r.Body, body)
				if body.Params.HasRemainQuota != "0" {
					wxPusher.Push(couponDetail.Title)
					//music.Play()
					log.Println(couponDetail.Title, "-----", body.Params.CouponQuota)
				} else {
					//log.Println(couponDetail.Title, "-----", body.Params.CouponQuota)
				}
			}
		})
		//https://content.95516.com/UCIWeb/outputWeb/coupon/coupon.html#/couponDetail?couponId=3102023031363076
		url := fmt.Sprintf("https://content.95516.com/koala-pre/koala/coupon/state?couponId=%s&cityCd=500000", couponDetail.CouponId)
		_ = c.Visit(url)
		time.Sleep(sleepTime)
		fmt.Println("----------", count, "----------")
	}

}

