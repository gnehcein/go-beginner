package main

import (
	"flag"
	"fmt"
)

var Input_pstrName = flag.String("name", "gerry", "input ur name")
var Input_piAge = flag.Int("age", 20, "input ur age")

func main()  {
	flag.Parse()
	arg:=flag.Args()
	fmt.Println(*Input_pstrName)
	fmt.Println(*Input_piAge)
	fmt.Printf("%T",arg)
}