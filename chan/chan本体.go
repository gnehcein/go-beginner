package main

import "fmt"

func main() {
	ch := make(chan int)
	ch1 := ch
	fmt.Println(ch, ch1)
}
