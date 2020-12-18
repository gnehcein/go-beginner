package main

import "fmt"

func main(){
	type dJG struct {
		Num 	int
		cha 	rune
		dd  	int
	}
	tu1:=dJG{
		Num :1,
		cha :'c',
		dd	:5,
	}
	p1:=&tu1
	p1.dd=2
	fmt.Println(tu1)

}
