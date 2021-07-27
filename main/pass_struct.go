package main

import "fmt"

type st struct {
	x int
}

func main() {
	thest, pt_st := f()
	pt_st.x = 5
	fmt.Println(thest, pt_st)
}

func f() (st, *st) {
	st := st{}
	return st, &st
}
