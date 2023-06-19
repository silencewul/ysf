package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	go test(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func test(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
