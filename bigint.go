package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x, y, z big.Int
	d := 5
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	x1, _ := x.SetString(s1, 10)
	y1, _ := y.SetString(s2, 10)
	z1 := z.Sub(x1, y1)
	fmt.Println(&d)
	fmt.Printf("%v,%v", &z, z1)
}
