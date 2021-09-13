package main

import "fmt"

type stru1 struct {
	m1 map[int]int
}

func main() {
	stru := new(stru1)
	fmt.Println(stru.m1 == nil) // 里面的引用属性没有被初始化
}
