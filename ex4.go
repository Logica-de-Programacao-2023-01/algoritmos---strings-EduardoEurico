package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Printf("Digite uma palvra: ")
	fmt.Scan(&x)
	s1 := x
	fmt.Println(strings.ToUpper(s1))
}
