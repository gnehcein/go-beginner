package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	ch1:=make(chan int)
	go func() {
		mu.Lock()
		<-ch1
		fmt.Println("f1get")
		time.Sleep(2*time.Second)
		fmt.Println("fffff")
	}()
	ch1<-1
	time.Sleep(1*time.Second)
	fmt.Println(1)
	time.Sleep(5*time.Second)
}
