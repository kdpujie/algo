package main

import (
	"fmt"
)

func main() {
	values := []int{1, 4, 3}
	BuildMaxHeap(len(values), values)
	fmt.Println(HeapPop(&values))
}

//向下调整(一个节点数据发生变化时, 会对其下所有子节点顺序产生影响)
func downAdjust(i, lenth int, array []int) {
	var left, right, largest = i*2 + 1, i*2 + 2, i
	fmt.Printf("\t ==maxHeapIfy: left=%d,right=%d,larget=%d,size=%d \n", left, right, largest, len(array))
	//fmt.Printf("\t ==maxHeapIfy: array[left]=%d,array[right]=%d,array[larget]=%d,size=%d \n",array[left],array[right],array[largest],len(array))
	if left <= lenth && array[largest] < array[left] {
		largest = left
	}
	if right <= lenth && array[largest] < array[right] {
		largest = right
	}
	if i != largest {
		array[largest], array[i] = array[i], array[largest]
		downAdjust(largest, lenth, array)
	}
}

//向上调整(节点插入时)
func upAdjust(index int, array []int) {
	if index > 0 {
		parentIndex := index / 2
		if array[parentIndex] < array[index] {
			array[parentIndex], array[index] = array[index], array[parentIndex]
			upAdjust(parentIndex, array)
		}
	}
}

//插入
func HeapInsert(number int, array *[]int) {

	*array = append(*array, number)
	upAdjust(len(*array)-1, *array)
}

//建堆
func BuildMaxHeap(length int, array []int) {
	for i := length / 2; i >= 1; i-- {
		downAdjust(i, length-1, array)
	}
}

//取出堆顶元素, 并调整堆
func HeapPop(array *[]int) int {
	max := (*array)[0]
	(*array)[0] = 0
	downAdjust(0, len(*array)-1, *array)
	*array = (*array)[:len(*array)-1]
	return max
}

//堆排序
func HeapSort(size int, array []int) {
	BuildMaxHeap(size, array)
	for i := size - 1; i >= 1; i-- {
		fmt.Printf("HeapSort: i=%d array=%v \n", i, array)
		array[0], array[i] = array[i], array[0] //堆顶元素和最后元素互换位置.
		fmt.Printf("\t HeapSort: i=%d array=%v \n", i, array)
		downAdjust(0, i-1, array) //除最后元素外, 重建堆.

		fmt.Printf("\t HeapSort: i=%d array=%v \n", i, array)
	}

}
