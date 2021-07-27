package main

func test(x int, x2 string, y ...int) {
	// ...type只能用于参数的最后
	for _, i := range y {
		println(i)
	}
}

func main() {
	test(1, "2", 3, 4, 5)
}
