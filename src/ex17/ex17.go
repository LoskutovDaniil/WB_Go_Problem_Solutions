package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int64{1, 3, 4, 5}
	to_search1 := 3
	if !binarySearch(array, int64(to_search1)) {
		fmt.Println("expected true, got false")
		return
	}

	to_search2 := 6
	if binarySearch(array, int64(to_search2)) {
		fmt.Println("expected false, got true")
		return
	}

	fmt.Println("All tests passed.")
}

func binarySearch(array []int64, to_search int64) bool {
	index := sort.Search(len(array), func(i int) bool {
		return array[i] >= to_search
	})
	return index < len(array) && array[index] == to_search
}
