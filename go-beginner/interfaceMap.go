package main

import "fmt"

func main()  {
	m1:=make(map[interface{}]int)	//any type's key
	ch1:=make(chan int)
	m1[ch1]=89
	m1[5]=6
	m1["1r"]=8
	fmt.Println(m1[5],m1["1r"])
	fmt.Println(m1[ch1])
}
