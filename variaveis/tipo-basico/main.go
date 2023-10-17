package main

import "fmt"

func main() {
	var inteiro int
	var flutuante float64
	var caractere rune
	var texto string
	var booleano bool

	inteiro = 1
	flutuante = 3.11
	caractere = 'A'
	texto = "Golang"
	booleano = true

	fmt.Printf("Inteiro: %d\n", inteiro)
	fmt.Printf("Flutuante: %.2f\n", flutuante)
	fmt.Printf("Caractere: %c\n", caractere)
	fmt.Printf("Texto: %s\n", texto)
	fmt.Printf("Booleano: %v\n", booleano)

}
