package main

import (
	"time"
	"fmt"
)

func main() {
	intCh := make(chan int)

	go func() {
		for i := 1; ; i++ {
			intCh <- i
			time.Sleep(time.Second) // Добавляем задержку для наглядности
		}
	}()
		
	// Читаем из канала
	go func() {
		for val := range intCh {
			fmt.Println("Value:", val)
		}
	}()

	N := 5
	time.Sleep(time.Duration(N) * time.Second)
	close(intCh)
	fmt.Println("The program has ended")
}
