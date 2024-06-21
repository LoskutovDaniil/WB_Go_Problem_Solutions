package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, channelData <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range channelData {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func main() {
	// флаги командной строки
	numWorkers := flag.Int("workers", 3, "number of workers to start")
	flag.Parse()

	channelData := make(chan int)
	var wg sync.WaitGroup

	// наспамить горутины для чтения из канала
	for i := 1; i <= *numWorkers; i++ {
		wg.Add(1)
		go worker(i, channelData, &wg)
	}

	// Обработка сигнала завершения (Ctrl+C)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Запускаем горутину для обработки сигналов
	go func() {
		<-signalChan
		close(channelData)
	}()

	// Постоянная запись данных в канал
	go func() {
		for i := 1; ; i++ {
			channelData <- i
			time.Sleep(time.Second) // Добавляем задержку для наглядности
		}
	}()

	// Ожидание завершения всех воркеров
	wg.Wait()
	fmt.Println("All workers completed")
}
