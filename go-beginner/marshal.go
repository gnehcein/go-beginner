package main

import (
	"encoding/json"
	"fmt"
)

type jg struct {
	name			string	`json:"name"`
	Num 			int		`json:"num"`
}
func main()  {
	jg1:=jg{
		name: "idd",
		//结构内变量必须是大写开头，否则无法转换成[]uint8中的数据
		//并且unmarshal也会丢失这个变量
		Num	: 5,
	}
	bt,err:=json.Marshal(jg1)
	if err!=nil{
		return
	}

	var jg2 jg
	fmt.Println("jg1",jg1)
	err=json.Unmarshal(bt,&jg2)
	fmt.Println("bt",string(bt))
	fmt.Println("jg2",jg2)
	fmt.Printf("%T",bt)

}
