package math

// Average returns the mean value in a slice
func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}

	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Min returns the smallest value in slice
func Min(xs []float64) float64 {
	min := 2e127 - 1
	for _, x := range xs {
		if min > x {
			min = x
		}
	}
	return min
}

// Max returns the biggest value in slice
func Max(xs []float64) float64 {
	max := -2e127 + 1
	for _, x := range xs {
		if max < x {
			max = x
		}
	}
	return max
}
