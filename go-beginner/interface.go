package main

import "fmt"

type n1 struct {

}
type in1 interface {
	do()
}
func (n11 n1) do(){
	fmt.Println("do")
}
func main()  {
	var n11=n1{}
	in1:=n11
	in1.do()
}
