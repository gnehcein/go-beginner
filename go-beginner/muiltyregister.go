package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type X1 struct {
	num int
}

type X2 struct {
	str string
}

func (X X1) Name1(Args int,Reply *int) error {
	fmt.Println("X.num")
	*Reply=5+Args
	return nil
}
func (X X2) Name2(args int,reply *int) error {	//参数不用大写，结构体类名和方法名必须要大写
	fmt.Println("X.str")
	*reply=10+args
	return nil
}
func main()  {
	go func() {
		lis,_:=net.Listen("tcp",":156")
		var x1 X1	//这个也不用大写
		var x2 X2
		rpc.Register(x1)	//可注册多个struct
		rpc.Register(x2)

		for  {
			conn,_:=lis.Accept()
			rpc.ServeConn(conn)		//rpc...
		}
	}()
	client,_:=rpc.Dial("tcp",":156")	//有rpc.dial和net.dial
	var x int
	client.Call("X1.Name1",1,&x)		//调用的是结构类名，而不是实例名
	fmt.Println(x)
	client.Call("X2.Name2",1,&x)
	fmt.Println(x)
}