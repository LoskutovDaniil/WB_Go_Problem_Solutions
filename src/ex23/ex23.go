package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 2
	fmt.Println("Start slice:", slice)

	slice = deleteElement(slice, index)

	fmt.Println("Finish slice:", slice)
}

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
