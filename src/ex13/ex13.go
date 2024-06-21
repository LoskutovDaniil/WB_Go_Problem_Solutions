package main

import "fmt"

func main () {
	num1 := 3
	num2 := 5

	num1, num2 = num2, num1

	fmt.Println("num1:",num1, "\nnum2:", num2)
}

// второй вариант
// package main

// import "fmt"

// func main() {
//     num1 := 3
//     num2 := 5

//     a = num1 ^ num2
//     num2 = num1 ^ num2
//     num1 = num1 ^ num2

//     fmt.Println("num1:", num1)
//     fmt.Println("num2:", num2)
// }
