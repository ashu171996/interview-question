package main

import "fmt"

func main() {
	Permutaion([]rune("abc"), func(a []rune) {
		fmt.Println("permutaion is :", string(a))
	})
}

func Permutaion(a []rune, f func([]rune)) {
	permutaion(a, f, 0)
}
func permutaion(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	permutaion(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permutaion(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
