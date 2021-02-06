package main

import "fmt"

func main() {
	s1:=[]int{1,2,3,4,5}
	s3:=[]rune{'a','b'}
	s2:=string(s3)
	fmt.Println(s1,s2,s3)
}
