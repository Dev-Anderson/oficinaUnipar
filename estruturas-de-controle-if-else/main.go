package main

import "fmt"

func main() {
	a := 1

	if a == 1 {
		fmt.Println("entrou dentro do if")
	} else if a == 2 {
		fmt.Println("entrou dentro do else if")
	} else {
		fmt.Println("entrou dentro do else")
	}
}
