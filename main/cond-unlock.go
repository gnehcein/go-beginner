package main

import (
	"fmt"
	"sync"
	"time"
)

func main() { //一个goroutine在两个时间段占有锁：
	//1.第一次lock，在wait之前。
	//2.wait被唤醒之后
	//
	cd := sync.Cond{L: &sync.Mutex{}}
	go func() {
		cd.L.Lock()
		cd.Wait()
		fmt.Println("1")
	}()
	go func() {
		cd.L.Lock()
		cd.Wait()
		fmt.Println("2")
	}()
	go func() {
		cd.L.Lock()
		cd.Wait()
		fmt.Println("3")
	}()
	time.Sleep(1 * time.Second)
	//cd.Signal()		//signal和broadcast会通知wait的goroutine获取锁。
	cd.Broadcast()

	time.Sleep(1 * time.Second)
	cd.L.Lock() //因为锁没被释放，所以deadlock

}
