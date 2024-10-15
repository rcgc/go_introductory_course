package main

import "fmt"

func main() {
	var arr1 = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	arr2 := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	arr3 := [10]float32{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}

	fmt.Printf("arr1: (%T) %v\n", arr1, arr1)

	fmt.Printf("arr2: (%T) ", arr2)
	for i:=0; i< len(arr2); i++ {
		fmt.Printf("%d ", arr2[i])
	}

	fmt.Printf("\narr3: (%T) ", arr3)
	for _, v := range arr3{
		fmt.Printf("%.1f ", v)
	}
}