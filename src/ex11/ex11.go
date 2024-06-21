package main

import "fmt"

func main() {
	var1 := []int{10, 20, 30}
	var2 := []int{15, 22, 30}

	varMap := make(map[int]bool)

	for _, num := range var1 {
		varMap[num] = true
	}

	var intersection []int

	for _, num := range var2 {
		if varMap[num] {
			intersection = append(intersection, num)
		}
	}

	fmt.Println(intersection)
}
