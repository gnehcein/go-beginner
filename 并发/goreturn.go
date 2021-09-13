package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	} //可能会输出 3,4
	time.Sleep(1 * time.Second)
}
