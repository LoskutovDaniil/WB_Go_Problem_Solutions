package main

import (
	"fmt"
	"strings"
)

func unique(str string) bool {
	m := make(map[rune]bool)
	for _, char := range strings.ToLower(str) {
		if m[char] {
			return false
		}
		m[char] = true
	}
	return true
}

func main() {
	strs := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, str := range strs {
		fmt.Printf("%s: %v\n", str, unique(str))
	}
}
