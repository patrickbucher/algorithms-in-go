package sorting

import (
	"cmp"
	"math/rand"
	"testing"
)

type Sort[T cmp.Ordered] func([]T)

func BenchmarkBubbleSort(b *testing.B) {
	for b.Loop() {
		BubbleSort([]int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for b.Loop() {
		InsertionSort([]int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0})
	}
}

func testWith[T cmp.Ordered](t *testing.T, f Sort[T], xs []T) {
	f(xs)
	for i := 0; i < len(xs)-1; i++ {
		if xs[i] > xs[i+1] {
			t.Errorf("sorting: %v at %d > %v at %d\n", xs[i], i, xs[i+1], i)
		}
	}
}

func testRandInt(t *testing.T, f Sort[int], n int) {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(n)
	}
	testWith(t, f, numbers)
}
