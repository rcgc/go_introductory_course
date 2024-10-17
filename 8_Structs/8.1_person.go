package main

import "fmt"

type Person struct{
	Name string
	Age int
	Height float32
	Weight float32
	isAlive bool
}

func main(){
	p := Person{
		Name: "Juan",
		Age: 33,
		Height: 1.66,
		Weight: 70.3,
		isAlive: true,
	}

	fmt.Printf("Person: (%T) %v\n", p, p)
	fmt.Printf("Person: (%T) %+v\n", p, p)
	fmt.Printf("Person: (%T) %#v\n", p, p)
	fmt.Printf("Name: (%T) %v", p.Name, p.Name)
}