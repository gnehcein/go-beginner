package main

import "fmt"

func main()  {
	s1:=[]int{1,2,3,4,5}
	s2:=s1[2:5]
	fmt.Println(s1,s2)
	fmt.Printf("%p,%p",&s1,&s2)
	s2=append(s2,6)
	fmt.Printf("%p,%p",&s1,&s2)
}