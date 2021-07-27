package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
			defer wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		fmt.Println("ok1")
	}()
	go func() {
		wg.Wait()
		fmt.Println("ok2") //同时结束，可以这样用
	}()
	time.Sleep(3 * time.Second)
}
