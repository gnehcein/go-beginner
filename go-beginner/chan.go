package main

import (
	"fmt"
	"time"
)

func main()  {
	ch1:=make(chan int,5)

	go getchan(&ch1)
	ch1<-5
	time.Sleep(5*time.Second)
	ch1<-7
	ch1<-7
	fmt.Println(len(ch1),cap(ch1))	//len即是通道内消息的个数
}
func getchan(ch1 *chan int)  {
	fmt.Printf("收到了,%v\n",<-(*ch1))
}
