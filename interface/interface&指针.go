package main

import "fmt"

type in4 interface{}
type stru1 struct {
	val int
}

func main() {
	stru := new(stru1)
	var inter1 in4 = stru //传的是指针
	fmt.Println(inter1)
	stru.val = 4
	fmt.Println(inter1)
	var inter2 in4 = inter1 //传的也是指针
	stru.val = 5
	fmt.Println(inter2)
}
