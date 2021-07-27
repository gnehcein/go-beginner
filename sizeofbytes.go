package main

import (
	"fmt"
	"unsafe"
)

type t struct {
	x int
	y int
}

func main() {
	s := make([]int32, 10)
	arr := [...]int32{1, 2, 3}
	copy(s, arr[:])
	fmt.Println(s)
	s2 := ([]byte)("dlsjfljsdlfdlsjfljadjfddddddsjfljadjfdddddddddddddd")

	t1 := t{
		x: 1,
	}
	fmt.Println(len(s2))
	s1 := [1001]t{} //通过把某类型封装成数组元素，可以计算
	fmt.Println("Size of []int32:", unsafe.Sizeof(s))
	fmt.Println("Size of [1000]int32:", unsafe.Sizeof(t1))
	fmt.Println("Real size of s:", unsafe.Sizeof(s1))
}
