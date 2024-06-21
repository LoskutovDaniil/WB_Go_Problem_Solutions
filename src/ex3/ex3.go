package main

import (
	"fmt"
	"sync"
)

func main() {
	x := [5]int{2, 4, 6, 8, 10}

	results := make(chan int, len(x)) // Создаем буферизованный канал

	var wg sync.WaitGroup
	wg.Add(len(x))

	for _, num := range x {
		go numSquare(num, results, &wg) // Запускаем горутины для вычисления квадратов чисел
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var sumRes int
	for res := range results {
		sumRes += res // Суммируем результаты из канала
	}

	fmt.Println("Сумма квадратов чисел:", sumRes)
}

func numSquare(num int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	results <- num * num // Отправляем результат в канал
}
