package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	i := 0
	var on sync.Once //只能在外部，否则不起作用
	for i < 10 {

		on.Do(func() {
			fmt.Println(1)
		})
		i++
	}
	time.Sleep(1 * time.Second)
}
