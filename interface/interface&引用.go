package main

import (
	"fmt"
)

type in3 interface {
}
type stru3 struct {
	int
}

func main() {
	m1 := make(map[int]int)
	var inter3 in3 = m1
	m1[1] = 1
	fmt.Println(inter3)
	var inter4 in3 = inter3
	fmt.Println(inter4)
	m1[2] = 2
	fmt.Println(inter3, inter4) //map传的是一个引用，但是实际上，引用的各种类型，数据本体都是一个struct，具体看帖子
}
