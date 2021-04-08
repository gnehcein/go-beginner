package main

import (
	"fmt"
	"reflect"
)


type T2 struct {
	X	int
	y 	int
}
func main() {
	//Part1
	//x:=5
	//p:=&x
	//val:=reflect.ValueOf(p)
	//val1:=val.Elem()
	//val1.SetInt(10)	//有各种set，通过reflect.Value来设置
	//fmt.Println(x)

	//Part2
	t:=&T2{1,2}	//必须是指针开始，才能最后赋值
	valueoft:=reflect.ValueOf(t)	//必须是value这条路才能改值，如果type，无法改也无法转变成value
	stru:=valueoft.Elem()
	n1:=stru.Field(0)	//像c++一样，必须先进入本体再调用其方法
	n1.SetInt(1)
	fmt.Println(t.X)
}
