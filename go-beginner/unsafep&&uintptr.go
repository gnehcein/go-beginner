package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var v1 =int8(10)
	ptr1:=&v1
	p1:=unsafe.Pointer(ptr1)
	p2:=(*int8)(p1)		//(type)(被转化的对象)两个括号
	p3:=(*string)(p1) //unsafe.Pointer可以转换为任意类型
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(ptr1)
	fmt.Println(uintptr(unsafe.Pointer(ptr1)))


}
