package main

import "fmt"

func main() {
	a := [7]int{2, 3, 5, 7, 11, 13, 17}

	var b []int = a[2:5]
	fmt.Println(b)

	c := make([]int, 4)
	c[0] = 12
	c[3] = 1
	fmt.Println("c", c)

	d := make([]int, 0, 5)
	fmt.Println("d", d)
	fmt.Println("tamanho de d: ", len(d))

	e := d[:2]
	fmt.Println("e", e)
	fmt.Println(len(e))
}
