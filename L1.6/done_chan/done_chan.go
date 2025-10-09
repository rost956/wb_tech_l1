package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("goroutine: got stop signal from channel")
				return
			default:
				fmt.Println("working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(done)
	time.Sleep(300 * time.Millisecond)
}
