package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go getchan(&ch1)
	ch1 <- 5
	time.Sleep(5 * time.Second)
}
func getchan(ch1 *chan int) {
	fmt.Printf("收到了,%v", <-(*ch1))
}
