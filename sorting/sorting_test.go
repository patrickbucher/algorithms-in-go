package sorting

import (
	"cmp"
	"math/rand"
	"slices"
	"testing"
)

type Sort[T cmp.Ordered] func([]T)

const bigTestSize = 10_000

func TestBubbleSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testWith(t, BubbleSort, numbers)
}

func TestInsertionSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testWith(t, InsertionSort, numbers)
}

func TestSelectionSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testWith(t, SelectionSort, numbers)
}

func TestQuickSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testWith(t, QuickSort, numbers)
}

func TestBubbleSortLarge(t *testing.T) {
	testRandInt(t, BubbleSort, bigTestSize)
}

func TestInsertionSortLarge(t *testing.T) {
	testRandInt(t, InsertionSort, bigTestSize)
}

func TestSelectionSortLarge(t *testing.T) {
	testRandInt(t, SelectionSort, bigTestSize)
}

func TestQuickSortLarge(t *testing.T) {
	testRandInt(t, QuickSort, bigTestSize)
}

func BenchmarkBubbleSortSmall(b *testing.B) {
	for b.Loop() {
		BubbleSort([]int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0})
	}
}

func BenchmarkInsertionSortSmall(b *testing.B) {
	for b.Loop() {
		InsertionSort([]int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0})
	}
}

func BenchmarkSelectionSortSmall(b *testing.B) {
	for b.Loop() {
		SelectionSort([]int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0})
	}
}

func BenchmarkQuickSortSmall(b *testing.B) {
	for b.Loop() {
		QuickSort([]int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0})
	}
}

func BenchmarkBubbleSortLarge(b *testing.B) {
	for b.Loop() {
		BubbleSort(randomNumbers(bigTestSize))
	}
}

func BenchmarkInsertionSortLarge(b *testing.B) {
	for b.Loop() {
		InsertionSort(randomNumbers(bigTestSize))
	}
}

func BenchmarkSelectionSortLarge(b *testing.B) {
	for b.Loop() {
		SelectionSort(randomNumbers(bigTestSize))
	}
}

func BenchmarkQuickSortLarge(b *testing.B) {
	for b.Loop() {
		QuickSort(randomNumbers(bigTestSize))
	}
}

func testWith[T cmp.Ordered](t *testing.T, f Sort[T], xs []T) {
	n := len(xs)
	ys := make([]T, n)
	copy(ys, xs)
	zs := slices.Sorted(slices.Values(xs))
	f(xs)
	if !slices.Equal(xs, zs) {
		t.Errorf("sorting %v: expected %v, got %v\n", ys, zs, xs)
	}
}

func testRandInt(t *testing.T, f Sort[int], n int) {
	testWith(t, f, randomNumbers(n))
}

func randomNumbers(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(n)
	}
	return numbers
}
