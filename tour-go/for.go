package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//SOBRE O LOOP
//um loop "for" que inicia uma variável "sum" com valor zero.
//Em seguida, ele executa o loop enquanto a variável "i" for menor que 10.
//A cada iteração do loop, o valor de "i"
//é somado ao valor de "sum". Após o loop, o valor final de "sum" é impresso na
//tela com "fmt.Println(sum)". O resultado será a soma dos números de 0 a 9, que é 45.
