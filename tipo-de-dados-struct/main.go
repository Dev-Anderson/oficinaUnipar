package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

type Trabalhador struct {
	Cargo string

	Pessoa
}

func main() {
	t := Trabalhador{"Mecanico", Pessoa{"Xico", 50}}
	fmt.Println(t)
}
