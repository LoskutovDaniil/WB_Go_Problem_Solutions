package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func Constructor(x, y float64) *Point {
	return &Point{x, y}
}

func Distance(var1, var2 *Point) float64 {
	dx := var2.x - var1.x
	dy := var2.y - var1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	var1 := Constructor(6, 8)
	var2 := Constructor(3, 4)

	fmt.Println("Calculated distance:", Distance(var1, var2))
}
