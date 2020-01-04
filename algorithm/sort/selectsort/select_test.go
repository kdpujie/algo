package selectsort

import "testing"

func TestSelectSort(t *testing.T) {
	values := []int{2, 1, 10, 5, 10, 3, 6, 10, 8, 5}
	selectSort(values)
	//t.Logf("sort after: %v", values)
	t.Errorf("sort after: %v", values)
}
