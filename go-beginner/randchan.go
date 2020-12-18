package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func GenerateintA()chan int {
ch := make(chan int, 10)
wg:=sync.WaitGroup{}
wg.Add(5)

	//启动 go routine 用于生成随机数，函数返回 个通道用于获取随机数
	go func() {
		for {
			ch <- rand.Int()
			fmt.Println("1")
			wg.Done()
			time.Sleep(1*time.Second)
			}
	}()
	wg.Wait()
	return ch
}

		func main() {
			ch := GenerateintA()
			fmt.Println(<-ch)
			fmt.Println(<-ch)
		}