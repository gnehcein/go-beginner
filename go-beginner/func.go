package main

import (
	"fmt"
	"time"
)

func main(){
	//可以x,y,z string    多个相同类型参数的写法，分配率

	go foo("1")
	go foo("2")
	go foo("3")
	foo1(foo)
	time.Sleep(5)
}
func foo(x string)  {
	fmt.Println(x)
}

//func(string)就够了
func foo1(f func(string))  {
	f("gg")

}
