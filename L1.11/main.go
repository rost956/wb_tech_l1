package main

import (
	"fmt"
)

func intersect(a, b []int) []int {
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}

	var result []int

	for _, v := range b {
		if m[v] {
			result = append(result, v)
			m[v] = false
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	C := intersect(A, B)

	fmt.Println("A =", A)
	fmt.Println("B =", B)
	fmt.Println("Пересечение =", C)
}
