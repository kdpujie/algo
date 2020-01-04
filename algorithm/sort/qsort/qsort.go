package qsort

func QuickSort(values []int) {
	partition(values, 0, len(values)-1)
}

//划分
func partition(values []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	base := values[left]
	for i < j {
		for values[j] >= base && i < j { //从数组右边开始遍历，一直找到小于base的数为止
			j--
		}
		for values[i] <= base && i < j { //从数组左边开始遍历，一直找到大于base的数为止
			i++
		}
		if i < j {
			values[i], values[j] = values[j], values[i]
		}
	}
	values[left], values[i] = values[i], values[left] // 基准数放到合适的位置(i的位置)
	//递归分治
	partition(values, left, i-1)  //继续处理左边的，这里是一个递归的过程
	partition(values, i+1, right) //继续处理右边的 ，这里是一个递归的过程
}
