package main

import (
	"fmt"
	"time"
	"ysf/ysf"
)

func main() {
	//ysf.UStrat()
	//ysf.Start()

	//test.StartG()

	ysf.TStart(1)

	//fibonacci

	//ch := make(chan int)
	//quit := make(chan int)
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-ch)
	//	}
	//	quit <- 0
	//}()
	//fibonacci(ch, quit)

	// worker
	//done := make(chan bool)
	//go worker(done)
	//<-done

	//var wg sync.WaitGroup
	//wg.Add(2) // 添加两个goroutine
	//
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("goroutine 1")
	//	fmt.Println(time.Now())
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("goroutine 2")
	//	fmt.Println(time.Now())
	//	fmt.Println(time.Now())
	//}()
	//
	//wg.Wait() // 等待所有goroutine执行完成
	//fmt.Println("main function")
	//fmt.Println(time.Now())

}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(1 * time.Second)
	fmt.Println("done")
	done <- true
}

func fibonacci(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			break

		}
	}
}
