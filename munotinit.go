package main

import (
	"../test1/te2"

	"fmt"
	"sync"
	"time"
)

type ct struct {
	mu sync.Mutex
}

func (ct1 *ct) fun() {
	ct1.mu.Lock()
	fmt.Println(1)
	ct1.mu.Unlock()
}
func main() {
	ct1 := ct{}

	for i := 0; i < 10; i++ {
		go ct1.fun()
	}
	testz.Bb()
	time.Sleep(5 * time.Second)
}
