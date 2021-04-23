package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var x int32
	go func() {
		for  {
			fmt.Println(atomic.LoadInt32(&x))
		}
	}()
	for  {
		atomic.StoreInt32(&x,100)
	}
}
