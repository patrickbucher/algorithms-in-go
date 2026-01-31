package sorting

import "cmp"

func BubbleSort[T cmp.Ordered](numbers []T) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if i == j {
				continue
			}
			if cmp.Less(numbers[j], numbers[j-1]) {
				Swap(numbers, j-1, j)
			}
		}
	}
}

func InsertionSort[T cmp.Ordered](numbers []T) {
	n := len(numbers)
	for i := 1; i < n; i++ {
		x := numbers[i]
		for j := 0; j < i; j++ {
			if cmp.Less(x, numbers[j]) {
				for k := i; k > j; k-- {
					Swap(numbers, k, k-1)
				}
				numbers[j] = x
				break
			}
		}
	}
}

func Swap[T cmp.Ordered](s []T, i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
