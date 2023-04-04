package main

import "fmt"

func main() {
	var x string
	var y string
	fmt.Printf("Digite uma palvra: ")
	fmt.Scan(&x)
	s1 := x
	fmt.Printf("Digite uma palavra: ")
	fmt.Scan(&y)
	s2 := y
	s3 := s1 + " " + s2
	fmt.Printf(s3)
}
