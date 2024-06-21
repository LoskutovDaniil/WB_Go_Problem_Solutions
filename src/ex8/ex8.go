package main

import (
	"flag"
	"fmt"
)

func main() {
	var num int64
	fmt.Printf("%064b\n", num)

	numWorkers := flag.Int("bit", 1, "change to 1")
	flag.Parse()

	num ^= 1 << *numWorkers
	fmt.Printf("%064b", num)
}
