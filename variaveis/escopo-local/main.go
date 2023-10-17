package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		y := 20
		fmt.Printf("x e %d e y e %d\n", x, y)
	}

	//y nao esta acessivel fora do escopo do if
	// fmt.Printf("O valor de y: %d\n", y)
	fmt.Printf("O valor de x: %d\n", x)
}
