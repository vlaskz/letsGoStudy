package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	v := Min([]float64{1, 2})
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestMax(t *testing.T) {
	v := Max([]float64{1, 2})
	if v != 2 {
		t.Error("Expected 2, got ", v)
	}
}
