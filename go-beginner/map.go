package main

import "fmt"

func main(){
	x:=make(map[string]int)
	x["dd"]=5
	x["dd"]=6
	//map一样
	fmt.Printf("%T\n%v\n",x,x)
	fmt.Printf("%T\n%p\n",&x,&x)
}