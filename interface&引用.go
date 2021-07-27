package main

import "fmt"

type in3 interface {
}

func main() {
	m1 := make(map[int]int)
	var inter3 in3 = m1
	m1[1] = 1
	fmt.Println(inter3)
	var inter4 in3 = inter3
	fmt.Println(inter4)
	m1[2] = 2
	fmt.Println(inter3, inter4)
}
