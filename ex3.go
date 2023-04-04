package main

import (
	"fmt"
)

func main() {
	var x, y string
	valor := 0
	fmt.Printf("digte uma palavra ")
	fmt.Scan(&x)
	fmt.Printf("Um caractere ")
	fmt.Scan(&y)

	for i := 0; i < len(x); i++ {
		if string(x[i]) == y {
			valor++
		}
	}
	fmt.Println("A letra ", y, " apareceu ", valor, " vezes em ", x)
}
