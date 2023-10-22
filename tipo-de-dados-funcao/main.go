package main

import "fmt"

func soma(x, y int) int {
	return x + y
}

func soma2(x, y int) {
	a := x + y
	fmt.Println(a)
}

func main() {
	resultado := soma(1, 2)
	fmt.Println(resultado)
	soma2(1, 2)
}
