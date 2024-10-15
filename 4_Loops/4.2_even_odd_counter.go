package main

import "fmt"

func main() {
	limit := 11

	even_counter := 0
	odd_counter := 0

	i := 1

	for i <= limit {
		if i%2 == 0 {
			even_counter++
		} else {
			odd_counter++
		}
		i++
	}

	fmt.Println("Even numbers:", even_counter)
	fmt.Println("Odd numbers:", odd_counter)
}