package main

import "fmt"

func makeSet(words []string) []string {
	m := make(map[string]bool)
	for _, w := range words {
		m[w] = true
	}

	var result []string
	for w := range m {
		result = append(result, w)
	}

	return result
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := makeSet(words)

	fmt.Println("words =", words)
	fmt.Println("set   =", set)
}
