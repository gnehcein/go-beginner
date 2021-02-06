package main

import (
	"fmt"
	"sync"
	"time"
)

type st struct {
	mu 		sync.Mutex
	n1 		int
}
func main() { //mu不会阻碍其他方法直接使用结构的变量
	st1:=st{}
	go func() {
		st1.mu.Lock()
		st1.n1=5
		fmt.Println(st1.n1)
		time.Sleep(10*time.Second)
		st1.mu.Unlock()
	}()
	time.Sleep(1*time.Second)
	go func() {
		st1.n1=7
		fmt.Println(st1.n1)
	}()
	time.Sleep(25*time.Second)
	fmt.Println(st1.n1)
}
