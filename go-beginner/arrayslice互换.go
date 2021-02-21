package main

import "fmt"

func main() {
	arr:=[...]int{1,2}
	var sli0 []int
	copy(sli0,arr[:])	//把数组提取切片，然后复制，反之也一样,用数组的切片模式去承接切片
	fmt.Println(sli0)
}
