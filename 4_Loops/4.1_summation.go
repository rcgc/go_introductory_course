package main

import "fmt"

func main(){
	result := 0
	limit := 10

	for i := 1; i<=limit; i++ {
		result = result + i
	}

	fmt.Println("Result:", result)
}