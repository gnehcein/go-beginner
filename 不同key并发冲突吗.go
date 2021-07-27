package main

import (
	"fmt"
	"time"
)

func main() {
	//如果都是读，不会
	//如果读和写，或者同时写，报错
	map1 := make(map[int]int)
	map1[1] = 1
	map1[2] = 2
	go func() {
		var x int
		for {
			x, _ = map1[1]
		}
		fmt.Println(x)
	}()
	go func() {
		var y int
		for {
			y, _ = map1[2]
		}
		fmt.Println(y)
	}()
	time.Sleep(3 * time.Second)
}
