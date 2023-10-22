package main

import "fmt"

func main() {
	msg := make(chan string)

	go func() { msg <- "ping" }()

	msg2 := <-msg
	fmt.Println(msg2)
}
