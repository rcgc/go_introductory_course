package main

import (
	"fmt"
	"time"
)

func hello_world(n int) {
	fmt.Println(n, "Hello world!")
}

func main(){
	for i:=0; i<5; i++ {
		go hello_world(i)
	}

	// Calling Sleep method 
    time.Sleep(3 * time.Second)
}