package searching

import (
	"testing"
)

const benchmarkSize = 1_000_000

var tests = []struct {
	haystack []int
	needle   int
	index    int
}{
	{[]int{}, 42, -1},
	{[]int{17}, 42, -1},
	{[]int{42}, 42, 0},
	{[]int{1, 2, 3}, 4, -1},
	{[]int{1, 2, 3, 4}, 4, 3},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 17, -1},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 0},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 9},
}

func TestLinearSearch(t *testing.T) {
	for _, test := range tests {
		if actual := LinearSearch(test.haystack, test.needle); actual != test.index {
			t.Errorf("LinearSearch(%v, %v): expected %d, was %d\n",
				test.haystack, test.needle, test.index, actual)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	for _, test := range tests {
		if actual := BinarySearch(test.haystack, test.needle); actual != test.index {
			t.Errorf("BinarySearch(%v, %v): expected %d, was %d\n",
				test.haystack, test.needle, test.index, actual)
		}
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	haystack := make([]int, benchmarkSize)
	for i := 0; i < benchmarkSize; i++ {
		haystack[i] = i
	}
	for b.Loop() {
		LinearSearch(haystack, benchmarkSize)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	haystack := make([]int, benchmarkSize)
	for i := 0; i < benchmarkSize; i++ {
		haystack[i] = i
	}
	for b.Loop() {
		BinarySearch(haystack, benchmarkSize)
	}
}
