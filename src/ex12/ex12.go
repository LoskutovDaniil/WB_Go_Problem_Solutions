package main

import "fmt"

func main() {
	varMap := make(map[string]bool)

	varSlice := []string{"cat", "cat", "dog", "cat", "tree"}

	var properSet []string

	for _, str := range varSlice {
		if !varMap[str] {
			properSet = append(properSet, str)
			varMap[str] = true
		}
	}
	fmt.Println(properSet)
}
