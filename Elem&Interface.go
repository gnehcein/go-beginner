package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var r io.Reader = os.Stdin
	p1 := reflect.ValueOf(&r)
	fmt.Println(p1.Elem().Type())
	fmt.Println(p1.Elem().Elem().Type())
	fmt.Println(reflect.ValueOf(r).Type())
	//接口作为左值，&的时候是本体;作为右值或作为函数参数（其实这也是个右值）的时候是承载值
}
