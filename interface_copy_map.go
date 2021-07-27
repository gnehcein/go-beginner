package main

import "fmt"

type emptyin interface{}

func main() {
	var in1 emptyin = make(map[int]int)
	var in2 emptyin = make(map[int]int)
	m1 := in1.(map[int]int)
	m1[5] = 6
	m2 := in1.(map[int]int)
	fmt.Println(m2)
	in2 = in1 //in1,2所承接的是map的引用
	fmt.Println(in2)
	m2[10] = 11
	fmt.Println(in2)
}
