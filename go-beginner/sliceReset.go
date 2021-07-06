package main

import "fmt"

func main() {
	slice1:=make([]int,5,10)
	slice1[1]=5
	slice1=slice1[:0]	//reset只是改变len，即可修改的长度变为0，cap还是10
	fmt.Println(slice1)
	fmt.Println(len(slice1),cap(slice1))
}
