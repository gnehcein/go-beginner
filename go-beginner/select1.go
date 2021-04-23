package main

import (
	"fmt"
	"time"
)

func main() {
	ch1:=make(chan int,100000)

	for  {
		select {
		case ch1<-1:	//不堵塞时，先考虑case，都堵塞则走default
			fmt.Println(1)
			time.Sleep(10*time.Millisecond)
		default:
			fmt.Println(2)
		}
	}

}
