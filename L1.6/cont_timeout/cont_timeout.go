package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Millisecond)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine:", ctx.Err())
				return
			default:
				fmt.Println("working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
}
