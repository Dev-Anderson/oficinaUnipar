// constantes
package main

import "fmt"

const pi = 3.14

func main() {
	raio := 5.0
	area := pi * (raio * raio)
	fmt.Printf("A area de um circulo com raio %.2f e %.2f\n", raio, area)
}
