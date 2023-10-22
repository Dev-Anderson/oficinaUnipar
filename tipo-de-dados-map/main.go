package main

import "fmt"

func main() {
	var m1 map[string]int
	var m2 = make(map[string]int)
	var m3 = map[string]int{"population": 5000}
	var m4 = m3

	var m5 map[string]string
	m5 = make(map[string]string)

	fmt.Println(m1, m2, m3, m4, m5)
}
