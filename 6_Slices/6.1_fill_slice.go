package main

import "fmt"

func main(){
	var slice []int

	fmt.Println("Before")
	fmt.Printf("slice: (%T) %v length=%d\n", slice, slice, len(slice))

	for i:=0; i<10; i++ {
		// Append works on nil slices
		slice = append(slice, i)
	}

	fmt.Println("\nAfter")
	fmt.Printf("slice: (%T) %v length=%d\n", slice, slice, len(slice))
}