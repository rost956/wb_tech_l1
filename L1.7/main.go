package main

//Программа реализует безопасную для конкуренции запись и чтение данных в map через синхронизацию

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			m[fmt.Sprintf("k%d", id)] = id // безопасная запись
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	mu.Lock()
	fmt.Println("len:", len(m)) //безопасное чтение
	for k, v := range m {
		fmt.Println(k, "=", v)
	}
	mu.Unlock()
}
