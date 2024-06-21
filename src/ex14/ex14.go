package main

import "fmt"

func main() {
	num := 36
	str := "WB"
	pravda := true
	varChnl := make(chan int, 1)
	varChnl <- 9

	determineType(num)
	determineType(str)
	determineType(pravda)
	determineType(varChnl)
}

func determineType(value interface{}) {
	switch result := value.(type) {
	case int:
		fmt.Println("Type entered: int", result)
	case string:
		fmt.Println("Type entered: string", result)
	case bool:
		fmt.Println("Type entered: bool", result)
	case chan int:
		res, ok := <-result
		if ok {
			fmt.Println("Type entered: chan int", res)
		} else {
			fmt.Println("Channel closed")
		}
	case chan string:
		res, ok := <-result
		if ok {
			fmt.Println("Type entered: chan string", res)
		} else {
			fmt.Println("Channel closed")
		}
	default:
		fmt.Println("Unknown type")
	}
}
