package main

import (
	"fmt"
	"unsafe"
)

type Z struct {
	i int
	k int
}

func main() {
	ch := make(chan Z)
	chemp := make(chan struct{}) //非常优秀，空间为0
	chint := make(chan int)
	close(ch)
	close(chemp)
	close(chint)
	x := <-ch
	y := <-chemp
	z := <-chint
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(y)) //内存占用为0，如果只是相当于没有值的传输信号的话，非常合适
	fmt.Println(unsafe.Sizeof(z))
}
