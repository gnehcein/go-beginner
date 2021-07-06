package main

import "fmt"

var str1 ="\\ldsjfldf"
func main() {
	str1=str1[:len(str1)-2]
	str1+="1"
	fmt.Println(str1)
}
