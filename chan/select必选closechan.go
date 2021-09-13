package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	closech := make(chan int)
	close(ch1)

	for i := 0; i < 10; i++ { // 是的，<-(close chan) 是通着的，即使是特殊情况，也是必选
		go func(i int) {
			select {
			case <-ch1:
				fmt.Println(i)
			case closech <- 1: // 堵塞，不选
				fmt.Println("err")
			}
		}(i)
	}
	time.Sleep(2 * time.Second)
}
