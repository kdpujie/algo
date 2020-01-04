package bubblesort

/**
@description
	顺序->大的往前冒泡
	从前往后冒泡
时间复杂度：O(n^2)
实现思路：
	1.以i的位置为基准，[len(values)-1-i,len(values)-1]为有序序列。（减少对比次数）
	2.设置flag，如果在冒泡的过程中，无数据交换，说明序列已经为有序序列，停止后续对比。
**/
func BubbleSort(values []int) {
	flag := true
	for i := 0; i < len(values); i++ {
		flag = true
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			} //end if
		} //end for j
		if flag { //等于true,说明剩下的序列已经排好序了
			break
		}
	} // end for i
}

/**
@description
	逆序->大的往前冒泡
	从后往前冒泡
时间复杂度：O(n^2)
实现思路：
	以i的位置为基准，[0,i]为有序序列。从后往前两两对比，把无序序列中最小的冒泡到 有序序列
**/
func BubbleSort3(values []int) {
	for i := 0; i < len(values); i++ {
		for j := len(values) - 1; j > i; j-- {
			if values[j] > values[j-1] {
				values[j], values[j-1] = values[j-1], values[j]
			}
		}
	}
}
