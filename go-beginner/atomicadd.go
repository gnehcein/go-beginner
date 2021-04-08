package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var x int32
	fmt.Println(atomic.AddInt32(&x,1))

}
