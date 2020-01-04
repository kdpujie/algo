package selectsort

/**
@description
	顺序排序
时间复杂度：O(n^2)
实现思路：
	1.以i的位置为基准，[0,i]为有序序列，(i,size-1]为无序序列。每次循环定位出无序序列中最小的数，和i位置交换。
**/
func selectSort(values []int) {
	for i := 0; i < len(values); i++ {
		minIndex := i
		for j := i + 1; j < len(values); j++ {
			if values[minIndex] > values[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			values[i], values[minIndex] = values[minIndex], values[i]
		}
	}
}
