package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ch1 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 1
	}()

	for i := 0; i < 5; i++ {
		fmt.Println(rand.Int63() % 20)
	}
	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-time.After(4 * time.Second):
		fmt.Println("4s")
	}

}
