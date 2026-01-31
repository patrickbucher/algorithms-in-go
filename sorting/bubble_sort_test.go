package sorting

import (
	"slices"
	"testing"
)

func TestBubbleSortSmall(t *testing.T) {
	numbers := []uint8{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	expected := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	BubbleSort(numbers)
	if !slices.Equal(numbers, expected) {
		t.Errorf("sorting numbers: expected %v, got %v\n", expected, numbers)
	}
}
