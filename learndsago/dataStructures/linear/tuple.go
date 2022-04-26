package linear

//powerSeries returns square and cube of an int
func powerSeries(n int) (int, int, error) {
	return n * n, n * n * n, nil
}
