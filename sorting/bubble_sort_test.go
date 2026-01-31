package sorting

import (
	"testing"
)

func TestBubbleSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	testWith(t, BubbleSort, numbers)
}

func TestBubbleSortLarge(t *testing.T) {
	testRandInt(t, BubbleSort, 1000)
}
