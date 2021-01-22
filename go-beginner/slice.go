package main

import "fmt"

func main(){
	s0:=[...]int{1,2,3,4,5}
	s1:=[]int{1,2,3,4}
	s2:=s1
	s2[0]=5
	s3:=s1
	x:=9
	//if s1==s2{    不能比较
	//}
	fmt.Printf("%T,,,%T,,,%T,,,,,%T\n",s0,s1,&s2,&x)
	//和c语言不同的是c数组名是指针，没有办法表示数组的实体，
	// 而golang中用数组/切片名表示实体，用& 数组/切片名表示指针
	fmt.Println(s1)
	fmt.Println(s1,s2)
	fmt.Print(s3)

}