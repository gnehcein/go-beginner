package main

import (
	"fmt"
	"time"
)

func main() {
	ch1:=make(chan int)

	go func(chan<- int) {

		time.Sleep(1*time.Second)


		ch1<-2

	}(ch1)
	fmt.Println(<-ch1)
	time.Sleep(3*time.Second)
}
