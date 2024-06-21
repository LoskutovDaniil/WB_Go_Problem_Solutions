package main

import (
	"fmt"
	"sync"
)

func main() {
	x := [3]int{2, 4, 6}

	numChannel := make(chan int, len(x))
	numSquareChannel := make(chan int, len(x))

	var wgChannelOne sync.WaitGroup
	wgChannelOne.Add(len(x))

	var wgChannelTwo sync.WaitGroup
	wgChannelTwo.Add(len(x))

	for _, num := range x {
		go numWrite(num, numChannel, &wgChannelOne) // Запускаем горутины
	}

	wgChannelOne.Wait()
	close(numChannel)

	for _, num := range x {
		go numSquareWrite(num, numSquareChannel, &wgChannelTwo) // Запускаем горутины
	}

	wgChannelTwo.Wait()
	close(numSquareChannel)

	for val := range numSquareChannel {
		fmt.Println("Value:", val)
	}
}

func numWrite(num int, results chan int, wgChannelOne *sync.WaitGroup) {
	defer wgChannelOne.Done()
	results <- num
}

func numSquareWrite(num int, results chan int, wgChannelTwo *sync.WaitGroup) {
	defer wgChannelTwo.Done()
	results <- num * 2
}
