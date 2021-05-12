package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string
	str="djafj"
	bt:=[]byte(str)
	p1:=(* string)(unsafe.Pointer(&bt))
	bt1:=[]byte{2,8}

	p2:=(* string)(unsafe.Pointer(&bt1))
	fmt.Println(*p1,*p2)
}
