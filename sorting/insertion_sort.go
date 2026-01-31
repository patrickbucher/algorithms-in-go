package sorting

import "cmp"

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
