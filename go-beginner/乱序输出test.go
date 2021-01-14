package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup
	for i:=0;i<5;i++ {
		time.Sleep(time.Second)//加上则不乱序，乱序是因为虽然goroutine生成是有顺序的，但println没有顺序
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
