package base

import (
	"testing"
)

var test = [3][6]int{
	{6, 5, 4, 3, 2, 1},
	{2, 3, 5, 7, 9, 11},
	{2, 5, 1, 7, 0, 1},
}

var res = [3][6]int{
	{1, 2, 3, 4, 5, 6},
	{2, 3, 5, 7, 9, 11},
	{0, 1, 1, 2, 5, 7},
}

func TestBubbleSorting(t *testing.T) {
	for i, _ := range test {
		array := test[i]
		bubbleSorting(array[:])
		if array != res[i] {
			t.Errorf("bubble Sorfting fail res: %v", array)
		}
	}
}
