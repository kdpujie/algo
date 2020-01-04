package qsort

import "testing"

func TestQuickSort(t *testing.T) {
	values := []int{2, 1, 10, 5, 10, 3, 6, 10, 8, 5}
	QuickSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("QuickSort() failed. Got", values, "Expected 1 2 3 4 5")
	}
}
