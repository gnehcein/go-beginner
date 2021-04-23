package main

import "fmt"

func main()  {
	ch:=make(chan int)
	defer func() {	//这样就不会报错panic了
		err:=recover()
		fmt.Println(err)
	}()
	close(ch)
	ch<-1
}
