package sequences

import (
	"slices"
	"testing"
)

const primesSmallTestSize = 100
const primesLargeTestSize = 100_000

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

var primeTests = []struct {
	l      int
	primes []int
}{
	{10, []int{2, 3, 5, 7}},
	{20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
	{30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
}

func TestPrimes(t *testing.T) {
	for _, test := range primeTests {
		expected := test.primes
		actual := Primes(test.l)
		if !slices.Equal(actual, expected) {
			t.Errorf("expected Primes(%d) to be %v, was %v\n", test.l, expected, actual)
		}
	}
}

func TestPrimeSieve(t *testing.T) {
	for _, test := range primeTests {
		expected := test.primes
		actual := PrimeSieve(test.l)
		if !slices.Equal(actual, expected) {
			t.Errorf("expected PrimeSieve(%d) to be %v, was %v\n", test.l, expected, actual)
		}
	}
}

func BenchmarkPrimesSmall(b *testing.B) {
	for b.Loop() {
		Primes(primesSmallTestSize)
	}
}

func BenchmarkPrimesLarge(b *testing.B) {
	for b.Loop() {
		Primes(primesLargeTestSize)
	}
}

func BenchmarkPrimeSieveSmall(b *testing.B) {
	for b.Loop() {
		PrimeSieve(primesSmallTestSize)
	}
}

func BenchmarkPrimeSieveLarge(b *testing.B) {
	for b.Loop() {
		PrimeSieve(primesLargeTestSize)
	}
}
