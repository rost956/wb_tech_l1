package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("defer runs before Goexit returns")
		fmt.Println("doing some work...")
		time.Sleep(300 * time.Millisecond)
		fmt.Println("calling runtime.Goexit()")
		runtime.Goexit()

	}()

	time.Sleep(1 * time.Second)
}
