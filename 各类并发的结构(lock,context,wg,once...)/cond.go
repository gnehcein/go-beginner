package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func main() {
	for i := 0; i < 10; i++ {
		go func(x int) { //为什么要设置参数？因为如果之后再传入i就是之后循环的了
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println(x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("传送")
	cond.Signal()
	time.Sleep(time.Second)
	fmt.Println("传送")
	cond.Signal()
	time.Sleep(time.Second)
	fmt.Println("传送")
	cond.Signal()
	time.Sleep(time.Second)
	fmt.Println("broadcast")
	cond.Broadcast()
	time.Sleep(2 * time.Second)

}
