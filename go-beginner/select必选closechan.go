package main

import (
	"fmt"
	"time"
)

func main() {
	ch1:=make(chan int)
	closech:=make(chan int)
	close(ch1)
	close(closech)
	for i:=0;i<100;i++{	//并不是。。随机的
		go func() {
			select {
			case <-ch1:
				fmt.Println("ok")
			case closech<-1:
				fmt.Println("err")
			}
		}()
	}
	time.Sleep(2*time.Second)
}
