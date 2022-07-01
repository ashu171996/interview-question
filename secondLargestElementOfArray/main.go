package main

import "fmt"

func main() {
	array := []int{4, 2, 1, 3, 5, 6}
	secondLargestElement(array)
}

func secondLargestElement(input []int) {
	max := input[0]
	min := 0
	for j := 1; j < len(input); j++ {
		if max < input[j] {
			min = max
			max = input[j]
		} else if min < input[j] {
			min = input[j]
		}
	}

	fmt.Println("second largest element", min)

}
