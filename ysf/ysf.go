package ysf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/gocolly/colly"
	"log"
	"os"
	"os/exec"
	"time"
)

type coupon struct {
	Title string
	CouponId string
}


type response struct {
	Resp string
	Msg string
	Params struct{
		CouponId string
		CouponQuota string
		CouponQuotaLow string
		HasRemainQuota string
		QuotaStatus string
		CouponQuotaPercent string
		CouponQuotaUseOut string
	}
}

var couponList = make([]coupon,0)

func Start()  {
	couponList = []coupon{
		{Title:"一綦向未来-满150元减50元通用消费",CouponId:"3102023042073905"},
		{Title:"一綦向未来-满200元减60元餐饮消费",CouponId:"3102023042073874"},
	}
	fmt.Println("请选择你要监控的卡券")

	for i,v := range couponList {
		fmt.Println(i+1,v.Title)
	}
	var num int
	_,err := fmt.Scanf("%d",&num)
	if err != nil || num < 1 || num > len(couponList) {
		log.Fatal("请输入1-",len(couponList),"数字")
	}
	var couponDetail = couponList[num-1]

	for {
		var buffer bytes.Buffer
			c := colly.NewCollector(
				// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
				colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
			)

			c.OnError(func(_ *colly.Response, err error) {
				log.Println("Something went wrong:", err)
			})

			c.OnResponse(func(r *colly.Response) {
				// 判断code
				if r.StatusCode == 200 {
					body := new(response)
					json.Unmarshal(r.Body,body)
					if body.Params.HasRemainQuota != "0" {
						//wxPusher.Push(couponDetail.Title)
						play()
						buffer.WriteString(couponDetail.Title+"  -----  "+body.Params.CouponQuota+"\n"+"\n")
						//log.Println(couponDetail.Title,"-----",body.Params.CouponQuota)
					} else {
						buffer.WriteString(couponDetail.Title+"  -----  "+body.Params.CouponQuota+"\n"+"\n")
						//fmt.Println(couponDetail.Title,"-----",body.Params.CouponQuota)
					}
				}
			})
	//https://content.95516.com/UCIWeb/outputWeb/coupon/coupon.html#/couponDetail?couponId=3102023031363076
			url := fmt.Sprintf("https://content.95516.com/koala-pre/koala/coupon/state?couponId=%s&cityCd=500000", couponDetail.CouponId)
			_ = c.Visit(url)
		fmt.Println(buffer.String())
		time.Sleep(2000*time.Millisecond)
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println("清除失败")
		}
		fmt.Println("---------------------------------")
	}


}

func play() {
	// 1. 打开mp3文件
	audioFile, err := os.Open("./1.mp3")
	if err != nil {
		log.Fatal(err)
	}
	// 使用defer防止文件描述服忘记关闭导致资源泄露
	defer audioFile.Close()

	// 对文件进行解码
	audioStreamer, format, err :=  mp3.Decode(audioFile)
	if err != nil {
		log.Fatal(err)
	}

	defer audioStreamer.Close()
	// SampleRate is the number of samples per second. 采样率
	// 通过采样率来更改播放速度, 只能通过整形倍数变换粒度太粗
	sr := format.SampleRate * 1
	_ = speaker.Init(sr, sr.N(time.Second / 10))

	// 重新采样 对然对电脑有信心还是先按照good进行测试，以免渲染不了
	// 从1开始升高音质，确保Mp3的本身是高音质的不然听着差别不大
	//var quality int = 6
	//resample := beep.Resample(quality, format.SampleRate, sr, audioStreamer)

	// 用于数据同步，当播放完毕的时候，回调函数中通过chan通知主goroutine
	done := make(chan bool)
	// 这里播放音乐
	speaker.Play(beep.Seq(audioStreamer, beep.Callback(func() {
		// 播放完成调用回调函数
		fmt.Println("111")
		done <- true
	})))

	// 增加控制信息
	for {
		select {
		// 等待音乐播放完成
		case <-done:
			return
		case <-time.After(time.Second):
			speaker.Lock()
			// 取的结果是每秒的SampleRate值，当前位置所在时间点
			//fmt.Println(format.SampleRate.D(audioStreamer.Position()).Round(time.Second))
			speaker.Unlock()
		}

	}
}
