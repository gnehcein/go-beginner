package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 2)
	fmt.Printf("%p\n", &slice1)
	slice1 = append(slice1, 1, 2, 3)
	fmt.Print("%p", &slice1)
	slice1 = append(slice1, 1)
	fmt.Print("%p\n", &slice1)
	slice1 = append(slice1, 1)
	fmt.Printf("%p\n", &slice1)

}
