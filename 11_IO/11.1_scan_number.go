package main

import "fmt"

func main() {
	var n int8 //[-127, 127]

	fmt.Println("Ingress an integer number [-128, 127]")
	fmt.Scan(&n)
	fmt.Printf("N: (%T) %v", n, n)
}