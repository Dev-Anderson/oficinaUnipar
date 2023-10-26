package main

import (
	"api/database"
	"fmt"
)

func main() {
	fmt.Println("API OK")
	database.Conectabanco()
}
