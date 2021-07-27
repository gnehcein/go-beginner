package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		close(ch1) //close之后就释放了
	}()
	msg, bo := <-ch1
	fmt.Println(msg, bo)
	msg, bo = <-ch1
	fmt.Println(msg, bo)
}
