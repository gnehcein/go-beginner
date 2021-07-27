package main

import (
	"fmt"
	"time"
)

func main() {
	x := makefunc()
	time.Sleep(5 * time.Second)
	x()
}

//竟然可以调用制造函数中的变量（即使制造函数已经结束了)
func makefunc() func() {
	t := 60
	return func() {
		fmt.Println(t, "finish")
	}

}
