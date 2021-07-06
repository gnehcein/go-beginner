package main

import "fmt"

type Flyer interface {
	Fly()
}
type Walker interface {
	Walk()
}
type Bird struct {
}
type Dog struct {
}

func (b Bird) Fly() {
	fmt.Println("the bird can fly!")
}
func (b Bird) Walk() {
	fmt.Println("the bird can walk!")
}
func (d *Dog) Walk() { //因为这里是用的指针传递，所以当调用接口方法时，必须要传递地址，不能是值。
	fmt.Println("the dog can walk!")
}
func main() {
	var i interface{} //声明一个接口对象
	i = new(Bird)     //创建一个Bird结构体变量，返回的是一个结构体指针，并赋值给接口
	v, b := i.(Flyer) //判断是否有飞行接口，
	if b {            //如判断成功，则v就是飞行接口类型
		v.Fly()
		fmt.Println(1)
	}
	v1, b := i.(Walker) //判断是否有行走接口
	if b {
		v1.Walk()
		fmt.Println(2)
	}
	//创建一个Dog结构体变量，然后把结构体地址赋值给接口，如果写成i=Dog{}，就无法调用Walk方法了，因为那里是接收的一个结构体指针。
	//所以将具体类型变量赋值给接口时，尽量使用指针传递，就是new(结构体类型)或者&结构体类型{}这两种写法，
	//这样不论类型变量的接口方法实现，是用值传递还是指针传递，都可以调用。
	i = &Dog{}
	v, b = i.(Flyer)
	if b {
		v.Fly()
		fmt.Println(3)
	}
	v1, b = i.(Walker)
	if b {
		v1.Walk()
		fmt.Println(4)
	}
}
