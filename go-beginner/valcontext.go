package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx0:=context.Background()
	ctx1,can:=context.WithCancel(ctx0)
	ctx2:=context.WithValue(ctx1,'a',1)
	ctx3:=context.WithValue(ctx2,'b',2)
	go gor(ctx3)
	time.Sleep(time.Second)
	can()
	time.Sleep(3*time.Second)

}
func gor(ctx1 context.Context){
	res:=ctx1.Value('b')
	fmt.Println(res)
	time.Sleep(2*time.Second)
	select {
	case <-ctx1.Done():
		fmt.Println("over")
	default:
		fmt.Println("denghui")
	}
	res2:=ctx1.Value('c')
	fmt.Println(res2)

}
