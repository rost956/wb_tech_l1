package main

// Программа реализует чтение из канала и запись в канал

import (
	"fmt"
	"time"
)

func main() {
	var N time.Duration = 5 //Время работы в секундах
	dataChannel := make(chan int)

	// Горутина пишущая в канал
	go func() {
		for i := 0; ; i++ {
			select {
			case <-time.After(N * time.Second):
				close(dataChannel)
				return
			default:
				dataChannel <- i
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Горутина читающая из канала
	go func() {
		for value := range dataChannel {
			fmt.Printf("Прочитано значение: %d\n", value)
		}
	}()

	// Ждем N секунд перед завершением
	<-time.After(N * time.Second)
}
