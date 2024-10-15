package main

import "fmt"

func main(){
	name := "Abraham"
	age := 1

	if age >= 18 {
		fmt.Println(name, "is", age, "years old. Therefore he's an adult.")
	}else {
		fmt.Println(name, "is", age, "years old. Therefore he's not an adult.")
	}
}