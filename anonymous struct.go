package main

import "fmt"

type stru11 struct {
	n1 int
	n2 int
}
type stru2 struct {
	stru11
}

func main() {
	intan := new(stru2)
	intan.n1 = 6
	intan.n2 = 7
	fmt.Println(*intan)
}
