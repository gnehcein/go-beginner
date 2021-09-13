package main

import (
	"fmt"
)

type empin interface{}
type VA struct {
	val int
}

func main() {
	V := VA{1}
	var in1 empin = V //!!!struct是复制的，接口承载的struct和V无关
	var in2 empin
	in2 = in1 //struct可以复制
	V.val = 2
	fmt.Println(in2)
	fmt.Println(in1)
}
