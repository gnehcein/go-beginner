package main

import (
	"fmt"
	"time"
)

func main() {

	//defer是串联的 go func 是并联的，return先返回然后defer，全执行完了再退出
	for t := 1; t < 5; t++ {
		defer fmt.Println("dd")
	}

	time.Sleep(1 * time.Second)
	return
}
