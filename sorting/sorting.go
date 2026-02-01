package sorting

import "cmp"

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

func Swap[T cmp.Ordered](s []T, i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
