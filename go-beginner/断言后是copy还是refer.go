package main

import "fmt"

type in5 interface {

}
type st struct {
	x  int
}
func main() {
	//是复制
	var i1 in5 = st{1}
	s1,ok := i1.(st)
	if ok {
		//...
	}
	s1.x = 10
	fmt.Println(i1,s1)
}
