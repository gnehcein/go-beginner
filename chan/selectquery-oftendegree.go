package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	bg := context.Background()
	bg1, can1 := context.WithCancel(bg)
	go func1(bg1)
	time.Sleep(1 * time.Second)
	can1()
	time.Sleep(2 * time.Second)
}
func func1(ctx context.Context) {
	fmt.Println("begin")
	i := 1
	for ; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("over")
			return
		default:
			time.Sleep(10 * time.Millisecond)
			fmt.Println(i)
		}
	}
}
