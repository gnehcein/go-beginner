package main

import (
	"fmt"
	"time"
)

func main() {
	d()
}
func d() {
	type x struct {
		x1 int
	}
	fir := x{1}
	fmt.Println(fir)
	go func() {
		sec := x{2}
		fmt.Println(sec)
	}()
	time.Sleep(1 * time.Second)
}
