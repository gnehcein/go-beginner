package main

import "fmt"

func main()  {
	//跟C不一样，指向数组/切片的指针是*[]int，而数组元素的指针是*int，C中是一样的
	sli1:=[...]int{1,2,3,4,5}
	p1:=&sli1
	fmt.Printf("%p\n",p1)
	sli2:=[]int{5,6,7}
	fmt.Printf("%T,\n %T\n",&sli1[0],&sli1)
	fmt.Println(sli2)
}
