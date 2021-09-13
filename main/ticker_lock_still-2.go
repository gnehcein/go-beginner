package main

import (
	"log"
	"time"
)

func main() {
	t := time.NewTicker(2 * time.Second)
	log.Println()
	time.Sleep(5 * time.Second)
	log.Println(1, <-t.C)
	log.Println(2, <-t.C)
	log.Println(3, <-t.C)

}
