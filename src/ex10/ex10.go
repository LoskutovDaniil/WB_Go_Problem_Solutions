package main

import "fmt"

func main() {
	tempMap := make(map[int][]float64)

	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	
	for _, num := range temp {
		key := int(num/10) * 10
		tempMap[key] = append(tempMap[key], num)
	}

	for key, temps := range tempMap {
        fmt.Printf("%d: %v\n", key, temps)
    }
}
