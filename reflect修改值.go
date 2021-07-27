package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	d := reflect.ValueOf(&x).Elem()
	fmt.Printf("%T\n", d)
	fmt.Println(d)
}
