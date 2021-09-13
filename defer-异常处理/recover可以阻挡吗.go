package main

import (
	"fmt"
	"time"
)

func main() {
	//可以阻挡！！！
	go func() {
		defer func() {
			err := recover()
			fmt.Println(err)
		}()
		ch := make(chan int)
		panic("ljlj") //情况一
		close(ch)
		ch <- 1 //情况二
	}()
	time.Sleep(2 * time.Second)
}
