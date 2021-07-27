package main

import "fmt"

const (
	x int = 6
	y     = 7
	z     = "lej"
	a     = iota
)

func main() {
	fmt.Printf("%T,%T,%T\n", x, y, z)
	fmt.Println(a)
	fmt.Println(1 >> 10)
}
