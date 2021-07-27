package main

import "fmt"

type G1 struct {
}

func (g G1) name() {
	fmt.Println("G1")
}
func (g *G1) name1() {
	fmt.Println("*G1")
}

func main() {
	var p = &G1{} //可以通过指针调用所指对象绑定的方法
	p.name()
	var g1 = G1{}
	g1.name1() //也可以通过struct调用指针绑定的方法
}
