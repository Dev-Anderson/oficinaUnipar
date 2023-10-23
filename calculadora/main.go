package main

import "fmt"

func Soma(a, b int) int {
	return a + b
}

func Subtracao(a, b int) int {
	return a - b
}

func Divisao(a, b int) int {
	return a / b
}

func Multiplicacao(a, b int) int {
	return a * b
}

func main() {
	soma := Soma(10, 10)
	fmt.Println("Soma = ", soma)
	subtracao := Subtracao(10, 5)
	fmt.Println("Subtracao = ", subtracao)
	divisao := Divisao(10, 2)
	fmt.Println("Divisao = ", divisao)
	multiplicacao := Multiplicacao(2, 2)
	fmt.Println("Multiplicacao = ", multiplicacao)
}
