package main

import "fmt"

type stru1 struct {
	m1  map[int]int
}

func main()  {
	stru:=new(stru1)	//都是初始化了
	fmt.Println(stru.m1[1])
}
