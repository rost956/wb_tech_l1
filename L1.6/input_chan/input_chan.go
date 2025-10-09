package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for v := range ch {
			fmt.Println("got:", v)
		}
		fmt.Println("goroutine: channel closed -> exit")
	}()

	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
	time.Sleep(300 * time.Millisecond)
}
