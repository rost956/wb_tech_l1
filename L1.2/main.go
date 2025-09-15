// Package main считает значения квадратов чисел из массива в горутинах
package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	fmt.Println(number * number)
}

func main() {
	mas := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, num := range mas {
		wg.Add(1)
		go square(&wg, num)
	}
	wg.Wait()
}
