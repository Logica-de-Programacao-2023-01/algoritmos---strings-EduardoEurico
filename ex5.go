package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Printf("Digite uma palvra: ")
	fmt.Scan(&x)
	////s1 := x
	var y int
	fmt.Printf("Escolha quantas letras se tornaram maiusculas: ")
	fmt.Scan(&y)

	for i := 0; i < y; i++ {
		fmt.Println(strings.ToUpper(string(x[i])))
	}
	for i := 0; i > i; i++ {
		fmt.Println(strings.ToUpper(string(x[i])))

	}
	fmt.Println(strings.ToUpper(string(x[:y] + x[y:])))
}
