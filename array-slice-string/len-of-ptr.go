package main

import "fmt"

func main() {
	sli := new([]int)
	fmt.Println(len(*sli)) //必须得*解引用，指针没有len方法
	fmt.Println(*sli == nil)
}
