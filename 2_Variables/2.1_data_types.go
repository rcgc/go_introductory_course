package main

import "fmt"

func main() {
	var name string
	name = "Juan"

	var age int8
	age = 33

	var height float32
	height = 1.66

	weight := 70.3

	var isAlive bool
	isAlive = true

	fmt.Println("Datos Personales:")
	fmt.Printf("Name: (%T) %v\n", name, name)
	fmt.Printf("Age: (%T) %v\n", age, age)
	fmt.Printf("Height: (%T) %v\n", height, height)
	fmt.Printf("Weight: (%T) %v\n", weight, weight)
	fmt.Printf("isAlive: (%T) %v\n", isAlive, isAlive)
}