package main

import "fmt"

type d struct {
	x int
	m map[int]int
}

func main() {
	m1 := make(map[int]int) //map slice chan，已初始化
	fmt.Printf("%p", &m1)
	m2 := new(d) //未初始化，返回地址，较可能分配在堆上
	fmt.Printf("%p", m2)
	var d1 d //未初始化，返回本值，较可能分配在栈上
	d1.m = make(map[int]int)
	d1.m[5] = 10
	fmt.Println(d1.m)
	//var ch chan int	//未初始化不能send或recv
	//go func() {ch<-1}()
	//<-ch
	var m3 map[int]int //未初始化，只可读，不可写
	fmt.Println(m3[5])
	m3 = map[int]int{}
}
