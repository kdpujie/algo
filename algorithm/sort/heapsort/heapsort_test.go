package main

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	values := []int{}
	//HeapSort(len(values),values)
	HeapInsert(2, &values)
	//HeapInsert(0,&values)
	//HeapInsert(2,&values)
	//HeapInsert(11,&values)
	t.Logf("after HeapInsert:%d \n", values)
	t.Logf("pop value=%d \n", HeapPop(&values))
	/*	t.Logf("pop value=%d \n",HeapPop(&values))
		t.Logf("pop value=%d \n",HeapPop(&values))*/
	t.Logf("after HeapPop:%d \n", values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("QuickSort() failed. Got", values, "Expected 1 2 3 4 5")
	}
}
