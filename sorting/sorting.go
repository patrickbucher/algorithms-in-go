package sorting

import (
	"cmp"
	"fmt"
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

func QuickSort[T cmp.Ordered](xs []T) {
	quickSort(xs, 0, len(xs)-1)
}

func quickSort[T cmp.Ordered](xs []T, lower, upper int) {
	fmt.Println("before", xs[lower:upper+1])
	n := upper - lower + 1
	if n <= 1 {
		return
	}
	pivot := xs[lower]
	p := lower
	i := lower + 1
	j := upper
	for i < j {
		for cmp.Compare(xs[i], pivot) <= 0 && i < j {
			i++
		}
		for cmp.Compare(xs[j], pivot) >= 0 && j > i {
			j--
		}
		if i < j {
			Swap(xs, i, j)
			p = i
		} else {
			p = j
		}
	}
	Swap(xs, lower, p)
	fmt.Println("after", xs[lower:upper+1])
	quickSort(xs, lower, p-1)
	quickSort(xs, p+1, upper)
}

func Swap[T cmp.Ordered](s []T, i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
