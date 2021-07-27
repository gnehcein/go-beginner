package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["dd"] = 5
	x["dd"] = 7
	//map一样
	//fmt.Printf("%T\n%v\n",x,x)
	//fmt.Printf("%T\n%p\n",&x,&x)
	x1 := make(map[int]string)
	x1[11] = "dkls"
	for k, v := range x1 {
		fmt.Println(k, v)
	}
}
