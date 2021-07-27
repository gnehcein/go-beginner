package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n1 := 5
	var p2 = unsafe.Pointer(&n1)
	fmt.Println(p2)
}
