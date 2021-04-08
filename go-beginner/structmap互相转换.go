package main

import (
	"encoding/json"
	"fmt"
)

type T1 struct {
	N  	int
	N2 	int
}
func main() {
	m1:=make(map[string]int)
	m2:=make(map[string]int)
	m1["N"]=4
	m1["N2"]=5
	m2["N3"]=4
	m2["N4"]=5
	t1:=T1{
		N:1,
		N2:2,
	}
	//js1, _ :=json.Marshal(t1)
	//err:=json.Unmarshal(js1,&m1)	//correct
	//fmt.Println(m1,err)
	//js2, _:=json.Marshal(m1)
	//err2:=json.Unmarshal(js2,&t1)	//correct
	//fmt.Println(t1,err2)
	//js3,_:=json.Marshal(m1)
	//err3:=json.Unmarshal(js3,&m2)	//转换后，N,N2,N3,N4都有
	//fmt.Println(m2,err3)
	js4,_:=json.Marshal(m2)
	err4:=json.Unmarshal(js4,&t1)	//t1中没有N3N4,所以转换不过来
	fmt.Println(t1,err4)
}
