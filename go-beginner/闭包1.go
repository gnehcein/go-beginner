package main

import "fmt"

func ret() func() {
	x:=1
	defer fmt.Println("over")
	return func() {
		x++
		fmt.Println(x)
	}
}
func main()  {
	f1:=ret()
	f1()
	f1()
}