package main

import "fmt"

func main() {
	var x string
	fmt.Printf("digite uma palavra ")
	fmt.Scan(&x)
	s := x
	fmt.Println("A palavra tem  ", len(s), " caracteres.")
}
