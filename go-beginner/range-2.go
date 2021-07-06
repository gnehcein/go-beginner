package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	res := []*int{}
	for _, v := range arr {
		//range实际上是用一个空间来依次拷贝 遍历到的东西
		res = append(res, &v)
	}
	//expect: 1 2 3
	fmt.Println(*res[0],*res[1],*res[2])
	//but output: 3 3 3
}
