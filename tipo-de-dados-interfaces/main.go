package main

import "fmt"

// Definindo uma interface chamada "Animal"
type Animal interface {
	Falar() string
}

// Definindo um tipo "Cachorro" que implementa a interface "Animal"
type Cachorro struct {
	Nome string
}

// Implementando o método "Falar" para o tipo "Cachorro"
func (c Cachorro) Falar() string {
	return "Woof! Meu nome é " + c.Nome
}

// Função que recebe qualquer tipo que implementa a interface "Animal"
func ApresentarAnimal(a Animal) {
	fmt.Println(a.Falar())
}

func main() {
	// Criando um cachorro
	meuCachorro := Cachorro{Nome: "Bob"}

	// Chamando a função ApresentarAnimal com o cachorro como argumento
	ApresentarAnimal(meuCachorro)
}
