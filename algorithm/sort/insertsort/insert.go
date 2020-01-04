package insertsort

/**
@description
	顺序排序
时间复杂度：O(n^2)
实现思路：
	1.通过比较找到合适的位置插入元素
	2.以i的位置为基准，[0,i]为有序序列，(i,size-1]为无序序列。
	3.从i(>0)开始遍历无序序列，找到有序序列中合适的位置，进行插入。
**/
func insertSort(values []int) {
	for i := 1; i < len(values); i++ {
		for j := i; j > 0; j-- {
			if values[j] >= values[j-1] {
				break
			} else {
				values[j], values[j-1] = values[j-1], values[j]
			}
		}
	}
}
