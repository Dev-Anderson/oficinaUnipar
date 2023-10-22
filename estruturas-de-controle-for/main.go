package main

import "fmt"

func main() {
	n := 5
	for n > 0 {
		n--
		println(n)
	}

	fmt.Println("Acabou o primeiro for")

	// declara i e acrescenta i em 1 a cada iteração
	for i := 0; i < 5; i++ {
		println(i)
	}

	fmt.Println("Acabou o segundo for")

	n = 5
	for i := 0; i < n; i++ {
		if i <= 2 {
			continue
		} else {
			println("i > 2 = ", i)
		}
	}

	fmt.Println("Acabou o terceiro for")

	n = 5
	for i := 0; i < n; i++ {
		if i == 2 {
			break
		} else {
			println("i: ", i)
		}
	}
}
