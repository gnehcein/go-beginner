package main

import (
	"fmt"
	"sync"
	"time"
)

func t(wg *sync.WaitGroup)  {//必须要传指针

	wg.Done()
	fmt.Println("free")
}
func main() {
	wg:=sync.WaitGroup{}
	wg.Add(1)
	go t(&wg)//waitgroup需要转换，cond context生成后直接传值
	time.Sleep(2*time.Second)
	wg.Wait()

}