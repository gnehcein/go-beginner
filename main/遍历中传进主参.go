package main

import (
	"fmt"
	"time"
)

type T1 struct {
	x int
}

func main() {
	//方法第一个参数（主参）跟后面的参数一样，类似于
	//go func(i int){
	//	...
	//}(i)
	sli := make([]*T1, 0, 5)
	for i := 0; i < 5; i++ {
		sli = append(sli, &T1{i})
	}
	for _, t := range sli {
		go t.f()
	}
	time.Sleep(1)
}
func (T1 *T1) f() {
	fmt.Println(T1.x)
}
