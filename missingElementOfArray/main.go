package main

import (
	"fmt"
)

func main() {
	a := []int{4, 5, 6, 1, 3}
	n := len(a) + 1
	missingElement(a, n)
}

func missingElement(a []int, n int) {
	sum := n * (n + 1) / 2
	for _, v := range a {
		sum -= v
	}
	fmt.Println("missing element", sum)
}
