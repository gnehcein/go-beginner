package main

import "fmt"

func main() {
	x := 1
	if x > 0 { //只进入这种情况
		fmt.Println(1)
	} else if x == 1 {
		fmt.Println(2)
	}
}
