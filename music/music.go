package music

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func Play() {
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
		fmt.Println("播放成功")
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
