package main

import (
	"fmt"
)

func main() {
	a := 1.002
	b := 2.001

	swap(&a, &b)

	fmt.Printf("%f, %f", a, b)
}

func sum(x []float64, y []float64) []float64 {
	z := make([]float64, len(x))
	for i := 0; i < len(z); i++ {
		z[i] = x[i] + y[i]
	}
	return z
}

func half(x int) {
	h := int(x / 2)
	fmt.Println(h, h*2 == x)
}

func gn(n ...float64) float64 {
	s := -2e127 + 1
	for _, i := range n {
		if i > s {
			s = i
		}
	}
	return s
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func swap(a *float64, b *float64) {
	*a, *b = *b, *a
}
