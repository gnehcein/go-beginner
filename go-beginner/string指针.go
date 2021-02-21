package main

import "fmt"

var str1="dlfj"
func main() {
	fmt.Println(ret(str1))
	s1:=[...]string{"你好"}
	s2:=[1]string{"你好"}
	fmt.Println(cap(s1), cap(s2),len(s1), len(s2))	//一样
}
func ret(str string) *string{
	return &str
}
