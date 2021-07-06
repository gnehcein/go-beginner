package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	for i:=0;i<5;i++{
		go func(i int) {
			time.Sleep(time.Duration(i) * time.Second)
			ch<-1
			fmt.Println(i)
		}(i)
	}
	time.Sleep(6*time.Second)
	for i:=0;i<5;i++{
		<-ch
	}
	time.Sleep(1*time.Second)
	fmt.Println("over")
}
