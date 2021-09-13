package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 1
	}()

	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-time.After(4 * time.Second): //即使有After,也会走default
		fmt.Println("4s")
		//default:
	}
}
