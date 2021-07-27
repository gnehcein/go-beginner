package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx1, cancel1 := context.WithCancel(ctx)
	ctx2, _ := context.WithCancel(ctx1)
	//go f1(ctx1)
	//go f2(ctx2)
	go f3(ctx2)
	time.Sleep(5 * time.Second)
	cancel1()

	time.Sleep(10 * time.Second)

}
func f1(ctx1 context.Context) {
	for {
		select {
		case <-ctx1.Done():
			defer fmt.Println("f1结束了")
			return
		default:
			time.Sleep(3 * time.Second)
			fmt.Println("f1工作")

		}
	}
}
func f2(ctx2 context.Context) {
	ctx := context.Background()
	go f2a(ctx)
	time.Sleep(2 * time.Second)
}
func f2a(ctx1 context.Context) {
	for {
		select {
		case <-ctx1.Done():
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("f2a工作了", time.Now())

		}
	}
}
func f3(ctx1 context.Context) {
	fmt.Println(22626)
	select {
	case <-ctx1.Done():
	default: //如果没有default，就只在一个case中等待
	}
	fmt.Println("finish")
}
