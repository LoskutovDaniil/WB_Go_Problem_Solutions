package main

import (
	"fmt"
	"time"
)

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	fmt.Println("Start sleeping...")
	Sleep(4)
	fmt.Println("Wake up!")
}
