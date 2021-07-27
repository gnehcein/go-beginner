package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("dd.txt", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString("111353***2**5")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}
