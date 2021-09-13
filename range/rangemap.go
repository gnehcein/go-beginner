package main

import "fmt"

func main() {
	//只是随机起点，多次固定顺序
	map1 := make(map[int]int)
	map1[1] = 1
	map1[4] = 4
	map1[2] = 2
	map1[3] = 3
	for k, v := range map1 {
		fmt.Println(k, v)
	}
}
