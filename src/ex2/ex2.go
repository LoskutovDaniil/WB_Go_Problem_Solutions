package main

import "fmt"

func main() {
	x := [5]int{2, 4, 6, 8, 10}

	results := make(chan int, len(x)) // Создаем буферизованный канал

	for _, num := range x {
		go numSquare(num, results) // Запускаем горутины
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(<-results) // Получаем и выводим результаты из канала
	}

}

func numSquare(num int, results chan int) {
	results <- num * num // Отправляем результат в канал
}
