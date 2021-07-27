package main

import "fmt"

func creat() func() {
	c := 0
	c++
	return func() {
		fmt.Println(c)
		c++
	}

}
func main() {
	d := creat()
	e := creat()
	d()
	e()
}
