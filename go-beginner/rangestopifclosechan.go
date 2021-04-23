package main

import (
	"fmt"
	"time"
)

func main()  {
	ch5:=make(chan int)
	go func(){
		for k:=range ch5{
			fmt.Println(k)
		}	//如果close，range一个chan会跳出循环。
		fmt.Println("over")
	}()
	for i:=0;i<10;i++ {
		ch5<-i
	}
	close(ch5)
	time.Sleep(time.Second)
}
