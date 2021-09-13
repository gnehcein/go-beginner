package main

import (
	"fmt"
	"reflect"
)

type st struct {
	x *int
}

func main() {
	var n1 = 1
	var n2 = 2
	st1 := st{&n1}
	st2 := st{&n2}
	fmt.Println(reflect.DeepEqual(st1, st2))
}
