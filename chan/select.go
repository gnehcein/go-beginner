package main

import (
	"fmt"
	"time"
)

func main() {
	for true {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println(1)
			break //无法突破循环，看同时亮起的是select
		case <-time.After(2 * time.Second):
			fmt.Println(2)
		}
		fmt.Println("seclct 后")
	}
	fmt.Println("over")
}
