package main

import "fmt"

type in interface {
	x(int)
	y(int)
}
type in2 interface {
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
 	var i1 in=T1{}	//接口对着sturct，不能接受struct的指针的方法，接口的视角把结构体和其指针实现的方法隔离了。
	var i2 in2 =&T1{}
	fmt.Println(i1,i2)
}
