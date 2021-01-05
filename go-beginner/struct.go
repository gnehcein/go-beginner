package main

import "fmt"

type dJG struct {
	Num 	int
	cha 	rune
	dd  	int
}
type fJG struct {
	Num 	int
	cha 	rune
	dd  	int
}
func (dJG)f1()  {
	fmt.Println("1ok")
}
func (fJG)f1(x int,y int)  {
	fmt.Println(x+y)
}
func main(){

	tu1:=dJG{
		Num :1,
		cha :'c',
		dd	:5,
	}
	tu2:=fJG{1, 'c', 2}
	tu2.f1(6,7)
	p1:=&tu1
	p1.dd=2
	fmt.Println(tu1)

}
