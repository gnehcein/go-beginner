package main

import "fmt"

type stru struct {
	x 		int
	y 		[]int
	z 		map[interface{}]int
}
func main() {
	t1:=stru{}
	t1.x=5
	t1.y=[]int{1,2,3}
	fmt.Println(t1.z==nil,t1.y==nil)
}
