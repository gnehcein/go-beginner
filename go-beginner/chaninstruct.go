package main

import (
	"fmt"
	"sync"
	"time"
)

type node struct {
	mu 		sync.Mutex
	ch1  	chan int
	d1  	int
	d2 		int
}
func (no *node)f1(){
	no.mu.Lock()
	no.d1++
	no.ch1<-1
	no.mu.Unlock()
}
func main() {
	node1:=node{
		ch1:make(chan int),
		d1:0,
	}//需要初始化ch1，否则是nil，传不了
	go (&node1).f1()
	x:=<-node1.ch1
	fmt.Println(node1.d1,x)

	time.Sleep(1*time.Second)
}
