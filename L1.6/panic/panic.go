package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("goroutine: recovered from panic:", r)
			}
		}()
		fmt.Println("working...")
		time.Sleep(300 * time.Millisecond)
		panic("stopping goroutine with panic")
	}()

	time.Sleep(1 * time.Second)
}
