package main

import "fmt"

func main()  {
	str:="ddd11"
	//str[0]='d' 不允许
	bytestr:=[]byte(str)
	bytestr[0]='n'
	str1:=string(bytestr)
	fmt.Println(str1)
}
