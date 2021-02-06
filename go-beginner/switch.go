package main

import "fmt"

func main()  {
	var x int
	switch x {
	case 1:
	case 0:
		fmt.Println(0)
		fallthrough
	case 2:fmt.Println(2)
	}
}
