package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[int]int)

	var mutex sync.Mutex

	for i := 0; i < 5; i++ {
		go func(n int) {
			// Захватываем мьютекс перед записью в мапу
			mutex.Lock()
			defer mutex.Unlock()

			// Записываем данные в мапу
			data[n] = n * n

			fmt.Printf("Added %d squared to map\n", n)
			fmt.Println("Current map state:", data)
		}(i)
	}

	// Ждем завершения работы всех горутин
	fmt.Println("Waiting for goroutines to finish...")
	// нажми энтер для завершения работы программы
	fmt.Scanln()
}
