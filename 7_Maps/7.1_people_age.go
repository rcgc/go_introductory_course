package main

import "fmt"

func main() {
	limit := 10
	var m1 = map[string]int{"Juan": 33, "John": 40, "Mary": 21, "Reyna": 29}
	m2 := map[string]string{"MX":"Mexico", "US":"United States", "GER":"Germany", "NZL":"New Zealand"}
	m3 := make(map[string]int)
	fmt.Printf("m1: (%T) %v\n", m1, m1)

	fmt.Printf("m2: (%T) ", m1)
	for i, v := range m2 {
		fmt.Printf("%s:%s ", i, v)
	}

	fmt.Println("\nBefore")
	fmt.Printf("m3: (%T) %v\n", m3, m3)
	for i := 1; i<=limit; i++ {
		if i % 2==0 {
			m3["even"]++
		} else {
			m3["odd"]++
		}
	}
	fmt.Println("After")
	fmt.Printf("m3: (%T) %v\n", m3, m3)
}