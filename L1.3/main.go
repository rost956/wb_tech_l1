// Package main реализует запись в канал и чтение данных из канала через воркеров
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Воркер %d: %d\n", id, data)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go <количество воркеров>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Ошибка: количество воркеров должно быть положительным числом")
		return
	}

	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	counter := 1
	for {
		ch <- counter
		counter++
		time.Sleep(100 * time.Millisecond)
	}
}
