package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("请输入用户名：")
	fmt.Scanf("我叫%s 今年%d岁", &name, &age) //expected space after %?
	fmt.Println(name, age)

}
