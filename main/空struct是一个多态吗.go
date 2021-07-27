package main

func main() {
	m12:=make(map[int]struct{})
	m12[1]= struct{
		x int,
	}{}	//不是多态
}
