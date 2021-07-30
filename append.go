package main

import "fmt"

type tt struct {
	x int
}

func main() {
	x1 := tt{x: 1}
	var s1 []tt
	s1 = append(s1, x1)
	s1 = append(s1, x1) //制造了一个新的struct，不是引用x1本体
	fmt.Println(s1)
	x1.x = 2
	fmt.Println(s1)
	s1 = append(s1, x1)
	fmt.Println(s1)
}
