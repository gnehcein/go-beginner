package main

import (
	"fmt"
	"unsafe"
)

type T11 struct {
	x1 	int
	x2  int
	p1  *int
	slice []int
}
func main() {
	num:=10
	t1:=T11{
		x1 : 0,
		x2 : 0,
		p1 : &num,
	}
	t1.slice=append(t1.slice, 1)
	t1.slice=append(t1.slice, 1)
	fmt.Println(unsafe.Sizeof(t1))
	
}
