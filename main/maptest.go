package main

import (
	"fmt"
)

func main() {
	var m map[int]int
	fmt.Println(m[0])
	assignMap(m)
	fmt.Println(m[0])
	m[0] = 0
	fmt.Println(m[0])
}

func assignMap(m map[int]int) {
	m = make(map[int]int)
	m[0] = 1
}
