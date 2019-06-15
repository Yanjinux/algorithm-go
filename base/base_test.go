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

func TestInsertsort(t *testing.T) {
	for i, _ := range test {
		array := test[i]
		Insertsort(array[:])
		if array != res[i] {
			t.Errorf("insert Sorfting fail res: %v", array)
		} else {

		}
	}
}

func TestSelectsort(t *testing.T) {
	for i, _ := range test {
		t.Logf(" src array:%v", test[i])
		array := test[i]
		Selectsort(array[:])
		if array != res[i] {
			t.Errorf("select Sorfting fail res: %v", array)
		} else {
			t.Logf(" res array:%v", array)
		}
	}
}

func TestQuicksort(t *testing.T) {
	for i, _ := range test {
		t.Logf(" src array:%v", test[i])
		array := test[i]
		Selectsort(array[:])
		if array != res[i] {
			t.Errorf("quick Sorfting fail res: %v", array)
		} else {
			t.Logf(" res array:%v", array)
		}
	}
}

func TestMergesort(t *testing.T) {
	for i, _ := range test {
		t.Logf(" src array:%v", test[i])
		array := test[i]
		Mergesort(array[:])
		if array != res[i] {
			t.Errorf("Mergesort Sorfting fail res: %v", array)
		} else {
			t.Logf(" res array:%v", array)
		}
	}
}

func TestStacksort(t *testing.T) {
	for i, _ := range test {
		t.Logf(" src array:%v", test[i])
		array := test[i]
		Stacksort(array[:])
		if array != res[i] {
			t.Errorf("Stacksort Sorfting fail res: %v", array)
		} else {
			t.Logf(" res array:%v", array)
		}
	}
}
