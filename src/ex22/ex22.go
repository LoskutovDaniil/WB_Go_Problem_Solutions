// Первый вариант
package main

import (
	"fmt"
	"math/big"
)

type data struct {
	operandA *big.Int
	operandB *big.Int
}

func (d data) sum() *big.Int {
	return new(big.Int).Add(d.operandA, d.operandB)
}

func (d data) sub() *big.Int {
	return new(big.Int).Sub(d.operandA, d.operandB)
}

func (d data) mul() *big.Int {
	return new(big.Int).Mul(d.operandA, d.operandB)
}

func (d data) div() *big.Int {
	if d.operandB.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("Error: division by zero")
		return nil
	}
	return new(big.Int).Div(d.operandA, d.operandB)
}

func main() {
	data := data{
		operandA: big.NewInt(13045020),
		operandB: big.NewInt(27430204),
	}

	min := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil)
	if data.operandA.Cmp(min) <= 0 || data.operandB.Cmp(min) <= 0 {
		fmt.Println("Error: operandA and operandB must be greater than 2^20")
		return
	}

	fmt.Println("sum:", data.sum())
	fmt.Println("Difference:", data.sub())
	fmt.Println("Product:", data.mul())
	if result := data.div(); result != nil {
		fmt.Println("Quotient:", result)
	}
}

// Второй вариант
// package main

// import "fmt"

// func main() {
// 	a := 1048577
// 	b := 1048577

// 	mul := a*b
// 	fmt.Println("a*b =", mul)

// 	div := a/b
// 	fmt.Println("a/b =", div)

// 	sum := a+b
// 	fmt.Println("a+b =", sum)

// 	sub := a-b
// 	fmt.Println("a-b =", sub)
// }
