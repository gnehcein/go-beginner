package main

func main() {
	// 不能返回字面量的地址
	// fmt.Println(&(makeint()))
	// fmt.Println(&(makestring()))
}
func makestring() string {
	return "dd"
}
func makeint() int {
	return 1
}
