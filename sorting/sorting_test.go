package sorting

import (
	"cmp"
	"math/rand"
	"slices"
	"testing"
)

type Sort[T cmp.Ordered] func([]T)

func TestBubbleSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testWith(t, BubbleSort, numbers)
}

func TestBubbleSortLarge(t *testing.T) {
	testRandInt(t, BubbleSort, 1000)
}

func TestInsertionSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testWith(t, InsertionSort, numbers)
}

func TestInsertionSortLarge(t *testing.T) {
	testRandInt(t, InsertionSort, 1000)
}

func TestSelectionSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testWith(t, SelectionSort, numbers)
}

func TestSelectionSortLarge(t *testing.T) {
	testRandInt(t, SelectionSort, 1000)
}

func TestQuickSortDebug(t *testing.T) {
	numbers := []int{7, 6}
	testWith(t, QuickSort, numbers)
}

func TestQuickSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testWith(t, QuickSort, numbers)
}

func TestQuickSortLarge(t *testing.T) {
	testRandInt(t, QuickSort, 100)
}

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

func BenchmarkSelectionSort(b *testing.B) {
	for b.Loop() {
		SelectionSort([]int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0})
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
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(n)
	}
	testWith(t, f, numbers)
}
