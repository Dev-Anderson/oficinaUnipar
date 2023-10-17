package main

import "fmt"

func main() {
	var numero [5]int
	numero[0] = 1
	numero[1] = 2

	fmt.Printf("Array: %v\n", numero)

	//slices (tamanho dinamico)
	numerosSlice := []int{1, 2, 3, 4, 5}
	numerosSlice = append(numerosSlice, 6)
	numerosSlice = append(numerosSlice, 7)

	fmt.Printf("Slice: %v\n", numerosSlice)
}
