package main

import (
	"fmt"
	"reflect"
)

type int1 int
type int2 int

func main() {
	//deepEqual比较type和value是否完全一致，层层递归

	//Array 	相同索引处的元素“深度”相等
	//Struct 	相应字段，包含导出和不导出，“深度”相等
	//Func 	只有两者都是 nil 时
	//Interface 	两者存储的具体值“深度”相等
	//Map 	1、都为 nil；2、非空、长度相等，指向同一个 map 实体对象，或者相应的 key 指向的 value “深度”相等
	//Pointer 	1、使用 == 比较的结果相等；2、指向的实体“深度”相等
	//Slice 	1、都为 nil；2、非空、长度相等，首元素指向同一个底层数组的相同元素，即 &x[0] == &y[0] 或者 相同索引处的元素“深度”相等
	//numbers, bools, strings, and channels,float32/64 	使用 == 比较的结果为真
	var n1 int1
	var n2 int2
	fmt.Println(n1, n2)
	fmt.Println(reflect.DeepEqual(n1, n2))
	var f1 float32 = 6
	var f2 float64 = 6
	fmt.Println(reflect.DeepEqual(f1, f2))
}
