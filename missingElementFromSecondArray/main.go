package main

import "fmt"

func main() {
	// given 2 array find missing element from array2
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{2, 3, 1, 0, 5}
	res := 0
	for _, v := range array1 {
		res = res ^ v
	}
	for _, v := range array2 {
		res = res ^ v
	}
	fmt.Println("missing element is", res)
}
