package main

import (
	"fmt"
	m "golang-book/chapter8/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	max := m.Max(xs)
	min := m.Min(xs)
	fmt.Println(avg)
	fmt.Println(max)
	fmt.Println(min)
}
