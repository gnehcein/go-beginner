package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f1, err := os.OpenFile("name1.txt", os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	data := make([]byte, 1024)
	var first, second string
	for {
		n, err := f1.Read(data)
		fmt.Println(n)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		for i := 0; i < n; i++ {
			if data[i] == '*' {
				first += "*"
			} else {
				second += string(data[i])
			}
		}
	}
	f2, err := os.OpenFile("name2.txt", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	f2.WriteString(first)
	f2.WriteString(second)

}
