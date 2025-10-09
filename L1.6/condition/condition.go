package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 1; ; i++ {
			fmt.Println("working...")
			if i == 5 {
				fmt.Println("goroutine: got stop by condition")
				return
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
}
