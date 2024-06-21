// Использование канала для сигнала завершения
package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				// Получили сигнал завершения, выходим из горутины
				fmt.Println("Goroutine stopped")
				return
			default:
				// Выполняем какую-то работу
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Даем горутине немного поработать
	time.Sleep(2 * time.Second)

	// Отправляем сигнал на остановку горутины
	stop <- true

	// Ждем, чтобы увидеть сообщение о завершении горутины
	time.Sleep(1 * time.Second)
}
