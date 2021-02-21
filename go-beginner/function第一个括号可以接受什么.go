package main

import "fmt"

type  T struct {
	x 	int
}
func main() {	//方法第一个括号只能接受一个参数，这个参数无论是地址还是对象本体，都可以传另一种
	var  t1	T	//而第二个括号的参数只能传相同类型
	t1.meth1(t1)	//error
	fmt.Println(t1.x)
	(&t1).meth2(&t1)	//error
}
func (t *T)meth1(t2  *T)  {
	t.x=5
}
func (t T)meth2(t2 	T)  {
	fmt.Println(t.x)
	t.x=6
}
