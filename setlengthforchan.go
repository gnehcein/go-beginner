package main

func main() {
	var ch1 chan [5]int
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- 1
		}
	}()
}
