package main

import (
	"fmt"
	"l1-24/point"
)

func main() {
	point1 := point.NewPoint(3, 4)
	point2 := point.NewPoint(7, 1)

	distance := point.Distance(point1, point2)
	fmt.Printf("Расстояние между точками: %.2f", distance)
}
