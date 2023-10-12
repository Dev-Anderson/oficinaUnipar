// variaveis em escopo local
package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		y := 20
		fmt.Printf("x: e %d, y e %d\n", x, y)
	}

	fmt.Printf("O valor de x: %d", x)
	// y nao esta acessivel fora do escopo do if
	// fmt.Printf("O valor de y: %d", y)
}
