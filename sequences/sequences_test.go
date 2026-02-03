package sequences

import (
	"slices"
	"testing"
)

func TestFibs(t *testing.T) {
	tests := []struct {
		n    int
		fibs []int
	}{
		{0, []int{}},
		{1, []int{1}},
		{2, []int{1, 1}},
		{3, []int{1, 1, 2}},
		{4, []int{1, 1, 2, 3}},
		{5, []int{1, 1, 2, 3, 5}},
		{6, []int{1, 1, 2, 3, 5, 8}},
		{7, []int{1, 1, 2, 3, 5, 8, 13}},
		{8, []int{1, 1, 2, 3, 5, 8, 13, 21}},
		{9, []int{1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{10, []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
	}
	for _, test := range tests {
		expected := test.fibs
		actual := Fibs(test.n)
		if !slices.Equal(actual, expected) {
			t.Errorf("expected Fibs(%d) to be %v, was %v\n", test.n, expected, actual)
		}
	}
}

func TestPrimes(t *testing.T) {
	tests := []struct {
		l      int
		primes []int
	}{
		{10, []int{2, 3, 5, 7}},
		{20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}
	for _, test := range tests {
		expected := test.primes
		actual := Primes(test.l)
		if !slices.Equal(actual, expected) {
			t.Errorf("expected Primes(%d) to be %v, was %v\n", test.l, expected, actual)
		}
	}
}
