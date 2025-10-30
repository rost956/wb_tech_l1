package main

import (
	"fmt"
	"math"
	"sort"
)

func bucket10(x float64) int {
	return int(math.Trunc(x/10.0)) * 10
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)
	for _, t := range temps {
		k := bucket10(t)
		groups[k] = append(groups[k], t)
	}

	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i, k := range keys {
		fmt.Printf("%d:{", k)
		for j, v := range groups[k] {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%.1f", v)
		}
		fmt.Print("}")
		if i < len(keys)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
