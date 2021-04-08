package main

type in interface {

}
type T struct {
	x1 	int
}

func (t *T)f1()  {

}
func main() {
	//t1:=&T{1}
	//var in1 in=t1
	//in1.f1()	报错，因为只能调用接口内的方法，in接口内没有方法
}
