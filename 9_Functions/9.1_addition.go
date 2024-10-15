package main

import "fmt"

func printMessage(){
	fmt.Println("Addition")
}

func addition(n1 int, n2 int) int {
	result := n1 + n2
	return result
}

func main(){
	n1 := 7
	n2 := 99

	printMessage()
	result := addition(7, 99)

	fmt.Println("Number 1 =", n1)
	fmt.Println("Number 2 =", n2)
	fmt.Println("Result =", result)
}