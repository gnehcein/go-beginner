package main

import (
	"context"
	"fmt"
)

func main() {

	for  {
		ctx:=context.Background()
		ctx1,can:=context.WithCancel(ctx)
		//放到这里就无限循环，因为有:=更新
		select {
		case <-ctx1.Done():
			fmt.Println("over")
			return
		default:
		}
		can()
	}
}
