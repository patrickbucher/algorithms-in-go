package sorting

import (
	"cmp"
	"math/rand"
	"slices"
	"testing"
)

type SortInPlace[T cmp.Ordered] func([]T)
type SortedCopy[T cmp.Ordered] func([]T) []T

const bigTestSize = 1000

func TestBubbleSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortInPlaceWith(t, BubbleSort, numbers)
}

func TestInsertionSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortInPlaceWith(t, InsertionSort, numbers)
}

func TestSelectionSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortInPlaceWith(t, SelectionSort, numbers)
}

func TestQuickSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortInPlaceWith(t, QuickSort, numbers)
}

func TestMergeSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortInPlaceWith(t, MergeSort, numbers)
}

func TestShellSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortInPlaceWith(t, ShellSort, numbers)
}

func TestTreeSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortInPlaceWith(t, TreeSort, numbers)
}

func TestBubbleSortedSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortedCopyWith(t, BubbleSorted, numbers)
}

func TestInsertionSortedSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortedCopyWith(t, InsertionSorted, numbers)
}

func TestSelectionSortedSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortedCopyWith(t, SelectionSorted, numbers)
}

func TestQuickSortedSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortedCopyWith(t, QuickSorted, numbers)
}

func TestMergeSortedSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortedCopyWith(t, MergeSorted, numbers)
}

func TestShellSortedSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortedCopyWith(t, ShellSorted, numbers)
}

func TestTreeSortedSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testSortedCopyWith(t, TreeSorted, numbers)
}

func TestBubbleSortLarge(t *testing.T) {
	testSortInPlaceRand(t, BubbleSort, bigTestSize)
}

func TestInsertionSortLarge(t *testing.T) {
	testSortInPlaceRand(t, InsertionSort, bigTestSize)
}

func TestSelectionSortLarge(t *testing.T) {
	testSortInPlaceRand(t, SelectionSort, bigTestSize)
}

func TestQuickSortLarge(t *testing.T) {
	testSortInPlaceRand(t, QuickSort, bigTestSize)
}

func TestMergeSortLarge(t *testing.T) {
	testSortInPlaceRand(t, MergeSort, bigTestSize)
}

func TestShellSortLarge(t *testing.T) {
	testSortInPlaceRand(t, ShellSort, bigTestSize)
}

func TestTreeSortLarge(t *testing.T) {
	testSortInPlaceRand(t, TreeSort, bigTestSize)
}

func TestBubbleSortedLarge(t *testing.T) {
	testSortedCopyRand(t, BubbleSorted, bigTestSize)
}

func TestInsertionSortedLarge(t *testing.T) {
	testSortedCopyRand(t, InsertionSorted, bigTestSize)
}

func TestSelectionSortedLarge(t *testing.T) {
	testSortedCopyRand(t, SelectionSorted, bigTestSize)
}

func TestQuickSortedLarge(t *testing.T) {
	testSortedCopyRand(t, QuickSorted, bigTestSize)
}

func TestMergeSortedLarge(t *testing.T) {
	testSortedCopyRand(t, MergeSorted, bigTestSize)
}

func TestShellSortedLarge(t *testing.T) {
	testSortedCopyRand(t, ShellSorted, bigTestSize)
}

func TestTreeSortedLarge(t *testing.T) {
	testSortedCopyRand(t, TreeSorted, bigTestSize)
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

func BenchmarkMergeSortSmall(b *testing.B) {
	for b.Loop() {
		MergeSort([]int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0})
	}
}

func BenchmarkShellSortSmall(b *testing.B) {
	for b.Loop() {
		ShellSort([]int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0})
	}
}

func BenchmarkTreeSortSmall(b *testing.B) {
	for b.Loop() {
		TreeSort([]int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0})
	}
}

func BenchmarkBubbleSortedSmall(b *testing.B) {
	numbers := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	for b.Loop() {
		BubbleSorted(numbers)
	}
}

func BenchmarkInsertionSortedSmall(b *testing.B) {
	numbers := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	for b.Loop() {
		InsertionSorted(numbers)
	}
}

func BenchmarkSelectionSortedSmall(b *testing.B) {
	numbers := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	for b.Loop() {
		SelectionSorted(numbers)
	}
}

func BenchmarkQuickSortedSmall(b *testing.B) {
	numbers := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	for b.Loop() {
		QuickSorted(numbers)
	}
}

func BenchmarkMergeSortedSmall(b *testing.B) {
	numbers := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	for b.Loop() {
		MergeSorted(numbers)
	}
}

func BenchmarkShellSortedSmall(b *testing.B) {
	numbers := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	for b.Loop() {
		ShellSorted(numbers)
	}
}

func BenchmarkTreeSortedSmall(b *testing.B) {
	numbers := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	for b.Loop() {
		TreeSorted(numbers)
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

func BenchmarkMergeSortLarge(b *testing.B) {
	for b.Loop() {
		MergeSort(randomNumbers(bigTestSize))
	}
}

func BenchmarkShellSortLarge(b *testing.B) {
	for b.Loop() {
		ShellSort(randomNumbers(bigTestSize))
	}
}

func BenchmarkTreeSortLarge(b *testing.B) {
	for b.Loop() {
		TreeSort(randomNumbers(bigTestSize))
	}
}

func BenchmarkBubbleSortedLarge(b *testing.B) {
	for b.Loop() {
		BubbleSorted(randomNumbers(bigTestSize))
	}
}

func BenchmarkSelectionSortedLarge(b *testing.B) {
	for b.Loop() {
		SelectionSorted(randomNumbers(bigTestSize))
	}
}

func BenchmarkInsertionSortedLarge(b *testing.B) {
	for b.Loop() {
		InsertionSorted(randomNumbers(bigTestSize))
	}
}

func BenchmarkQuickSortedLarge(b *testing.B) {
	for b.Loop() {
		QuickSorted(randomNumbers(bigTestSize))
	}
}

func BenchmarkMergeSortedLarge(b *testing.B) {
	for b.Loop() {
		MergeSorted(randomNumbers(bigTestSize))
	}
}

func BenchmarkShellSortedLarge(b *testing.B) {
	for b.Loop() {
		ShellSorted(randomNumbers(bigTestSize))
	}
}

func BenchmarkTreeSortedLarge(b *testing.B) {
	for b.Loop() {
		TreeSorted(randomNumbers(bigTestSize))
	}
}

func testSortInPlaceWith[T cmp.Ordered](t *testing.T, f SortInPlace[T], xs []T) {
	n := len(xs)
	ys := make([]T, n)
	copy(ys, xs)
	zs := slices.Sorted(slices.Values(xs))
	f(xs)
	if !slices.Equal(xs, zs) {
		t.Errorf("sorting %v:\nexpected\t%v\ngot\t\t%v\n", ys, zs, xs)
	}
}

func testSortedCopyWith[T cmp.Ordered](t *testing.T, f SortedCopy[T], xs []T) {
	ys := f(xs)
	zs := slices.Sorted(slices.Values(xs))
	if !slices.Equal(ys, zs) {
		t.Errorf("sorting %v:\nexpected\t%v\ngot\t\t%v\n", ys, zs, xs)
	}
}

func testSortInPlaceRand(t *testing.T, f SortInPlace[int], n int) {
	testSortInPlaceWith(t, f, randomNumbers(n))
}

func testSortedCopyRand(t *testing.T, f SortedCopy[int], n int) {
	testSortedCopyWith(t, f, randomNumbers(n))
}

func randomNumbers(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(n)
	}
	return numbers
}
