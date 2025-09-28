package main

/*
Контекст позволяет уведомить все горутины о необходимости завершения.
signal.Notify перехватывает SIGINT и отправляет в канал signalChannel.
sync.WaitGroup дает завершиться воркерам перед завершением основной горутины.
*/
import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d завершает работу\n", id)
			return
		default:
			fmt.Printf("Worker %d выполняет задачу\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	numWorkers := 3

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT)

	<-signalChannel
	fmt.Println("\nПолучен сигнал Ctrl+C. Завершаем работу...")

	cancel()
	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}
