package main

import (
	"bytes"
	"fmt"
)

func main() {
	bf:=bytes.Buffer{}
	bf.WriteByte('d')	//写字符
	bf.WriteString("sdf")
	fmt.Println(bf.String())	//提取，不是抽取
	fmt.Println(bf.String())
}
