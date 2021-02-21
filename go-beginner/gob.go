package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type ta struct {
	X 	int
	Y 	int
	M 	map[int]interface{}
}
func main() {
	var t1,t2,t3,t4,t5 ta
	var num1 	int
	var bf1 	bytes.Buffer
	enc:=gob.NewEncoder(&bf1) 	//开启一个输送接口
	t1.M=make(map[int]interface{})
	t1.M[1]=50
	t1.X=5
	enc.Encode(t1)

	t2.Y=8
	enc.Encode(t2)		//在bf1中按照顺序存放

	enc.Encode(5)

	dec:=gob.NewDecoder(&bf1)
	dec.Decode(&t4)		//按照顺序解码
	dec.Decode(&t3)
	dec.Decode(&t5)		//对应着encode（5)
	dec.Decode(&num1)	//已经被上一个消耗掉了,所以是0


	fmt.Println(t3)
	fmt.Println(t4)
	fmt.Println(t5)
	fmt.Println(num1)
}
