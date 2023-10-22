package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("Tamanho de int em bytes: %d\n", unsafe.Sizeof(int(0)))
}
