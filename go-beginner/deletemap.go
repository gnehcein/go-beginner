package main

import (
	"fmt"
	"time"
)

func main() {
	m1:=make(map[int]int)
	m1[1]=1145
	m1[2]=3543636
	m1[3]=466536

	delete(m1,1)
	delete(m1,1)
	delete(m1,1)
	delete(m1,1)
	delete(m1,1)	//没问题
	time.Sleep(1*time.Second)
	val,exist:=m1[1]
	fmt.Println(val,exist)
}
