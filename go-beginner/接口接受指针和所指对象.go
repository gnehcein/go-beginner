package main

import "fmt"

type in21 interface {
	x(int)
	y(int)
}
type in22 interface {
	z(int)
}

type T1 struct {}

func (t T1) x(int)  {	//struct方法

}
func (t T1) y(int)  {	//struct方法

}

func (t *T1) z(int)  {	//指针方法

}
func main() {
 	var i1 in21=&T1{}	//指针带着struct和指针本身的方法
	var i2 in22 =T1{}	//报错，因为结构体只带着自身的方法
	fmt.Println(i1,i2)
}
