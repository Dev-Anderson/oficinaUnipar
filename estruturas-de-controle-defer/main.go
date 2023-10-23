package main

import "fmt"

func main() {
	// Será executado no final da função main()
	defer fmt.Println("Golang e legal!")
	defer fmt.Println("World")

	fmt.Println("Hello")
}
