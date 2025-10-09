package main

import (
	"fmt"
	"time"
)

func main() {
	stopAfter := time.After(1200 * time.Millisecond)

	go func() {
		for {
			select {
			case <-stopAfter:
				fmt.Println("goroutine: time.After fired -> exit")
				return
			default:
				fmt.Println("working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
}
