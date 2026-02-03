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
	n := len(xs)
	if n == 0 {
		return []T{}
	}
	if n == 1 {
		return []T{xs[0]}
	}
	pivot := xs[0]
	smaller, larger := make([]T, 0), make([]T, 0)
	for i := 1; i < n; i++ {
		if cmp.Less(xs[i], pivot) {
			smaller = append(smaller, xs[i])
		} else {
			larger = append(larger, xs[i])
		}
	}
	ys := make([]T, 0, n)
	ys = append(ys, QuickSorted(smaller)...)
	ys = append(ys, pivot)
	ys = append(ys, QuickSorted(larger)...)
	return ys
}

func MergeSort[T cmp.Ordered](xs []T) {
	n := len(xs)
	if n <= 1 {
		return
	}
	p := n / 2
	q := n - p
	ys, zs := make([]T, p), make([]T, q)
	for i := 0; i < p; i++ {
		ys[i] = xs[i]
	}
	for i := p; i < n; i++ {
		zs[i-p] = xs[i]
	}
	MergeSort(ys)
	MergeSort(zs)
	var v T
	l := 0
	r := 0
	for i := 0; i < n; i++ {
		if l >= p {
			v = zs[r]
			r++
		} else if r >= q {
			v = ys[l]
			l++
		} else if cmp.Less(ys[l], zs[r]) {
			v = ys[l]
			l++
		} else {
			v = zs[r]
			r++
		}
		xs[i] = v
	}
}

func MergeSorted[T cmp.Ordered](xs []T) []T {
	n := len(xs)
	if n == 1 {
		return []T{xs[0]}
	}
	p := n / 2
	q := n - p
	ys, zs := MergeSorted(xs[0:p]), MergeSorted(xs[p:n])
	var v T
	l := 0
	r := 0
	vs := make([]T, n)
	for i := 0; i < n; i++ {
		if l >= p {
			v = zs[r]
			r++
		} else if r >= q {
			v = ys[l]
			l++
		} else if cmp.Less(ys[l], zs[r]) {
			v = ys[l]
			l++
		} else {
			v = zs[r]
			r++
		}
		vs[i] = v
	}
	return vs
}

func ShellSort[T cmp.Ordered](xs []T) {
	n := len(xs)
	fibs := Fibs(n)
	gaps := make([]int, 0)
	for i := 0; i < n; i++ {
		if fibs[i] >= n {
			break
		}
		gaps = append(gaps, fibs[i])
	}
	// FIXME: Fibonacci gaps are too big; use Prime Numbers instead
	for g := len(gaps) - 2; /* skip duplicate 1 */ g >= 0; g-- {
		d := gaps[g]
		for i := 0; i+d < n; i++ {
			if cmp.Less(xs[i+d], xs[i]) {
				Swap(xs, i, i+d)
			}
		}
	}
}

func Swap[T cmp.Ordered](s []T, i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

func Fibs(n int) []int {
	if n <= 0 {
		return []int{}
	}
	a, b := 1, 1
	fibs := make([]int, n)
	for i := 0; i < n; i++ {
		fibs[i] = a
		c := a + b
		a = b
		b = c
	}
	return fibs
}
