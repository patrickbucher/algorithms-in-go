package sorting

import (
	"cmp"
)

func BubbleSort[T cmp.Ordered](xs []T) {
	n := len(xs)
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if i == j {
				continue
			}
			if cmp.Less(xs[j], xs[j-1]) {
				Swap(xs, j-1, j)
			}
		}
	}
}

func BubbleSorted[T cmp.Ordered](xs []T) []T {
	ys := make([]T, len(xs))
	copy(ys, xs)
	BubbleSort(ys)
	return ys
}

func InsertionSort[T cmp.Ordered](xs []T) {
	n := len(xs)
	for i := 1; i < n; i++ {
		x := xs[i]
		for j := 0; j < i; j++ {
			if cmp.Less(x, xs[j]) {
				for k := i; k > j; k-- {
					Swap(xs, k, k-1)
				}
				xs[j] = x
				break
			}
		}
	}
}

func InsertionSorted[T cmp.Ordered](xs []T) []T {
	ys := make([]T, len(xs))
	copy(ys, xs)
	InsertionSort(ys)
	return ys
}

func SelectionSort[T cmp.Ordered](xs []T) {
	n := len(xs)
	for i := 0; i < n; i++ {
		min, m := xs[i], i
		for j := i + 1; j < n; j++ {
			if cmp.Less(xs[j], min) {
				min = xs[j]
				m = j
			}
		}
		Swap(xs, i, m)
	}
}

func SelectionSorted[T cmp.Ordered](xs []T) []T {
	// TODO: with a copy, there's a simpler implementation possible
	ys := make([]T, len(xs))
	copy(ys, xs)
	SelectionSort(ys)
	return ys
}

func QuickSort[T cmp.Ordered](xs []T) {
	quickSort(xs, 0, len(xs)-1)
}

func quickSort[T cmp.Ordered](xs []T, lower, upper int) {
	n := upper - lower + 1
	if n <= 1 {
		return
	}
	pivot := xs[lower]
	i := lower + 1
	j := upper
	for i < j {
		for i <= j && cmp.Compare(xs[i], pivot) <= 0 {
			i++
		}
		for j >= i && cmp.Compare(xs[j], pivot) >= 0 {
			j--
		}
		if i < j {
			Swap(xs, i, j)
		}
	}
	p := upper
	for cmp.Compare(xs[p], pivot) >= 0 && p > lower {
		p--
	}
	Swap(xs, lower, p)
	quickSort(xs, lower, p-1)
	quickSort(xs, p+1, upper)
}

func QuickSorted[T cmp.Ordered](xs []T) []T {
	// TODO: with a copy, there's a simpler implementation possible
	ys := make([]T, len(xs))
	copy(ys, xs)
	QuickSort(ys)
	return ys
}

func Swap[T cmp.Ordered](s []T, i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
