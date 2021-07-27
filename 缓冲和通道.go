package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	go func() {
		x := <-ch1
		fmt.Println(x)
	}()
	ch1 <- 5
	time.Sleep(5 * time.Second)
}
