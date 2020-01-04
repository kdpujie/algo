package main

import "log"

func main() {
	mm := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	log.Printf("是否存在：%v", banaryMatrix(mm, 16))
}

/**
1.问题描述
	在一个二维数组中，每一行都按升序排列，每列也都按升序排列。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
	1 2 8  9
	2 4 9  12
	4 7 10 13
	6 8 11 15
2.思路：
	从右上角开始对比过程，如果目标数字比9大，则剔除当前行；如果目标数字比9小，则剔除当前列。
	比如寻找数字7：7小于9，所以剔除9所在的整列， 下一次比较8的位置。
**/

func banaryMatrix(mm [][]int, numbers int) bool {
	if len(mm) < 1 {
		return false
	}
	var rowIndex, columIndex = 0, len(mm[0]) - 1
	for rowIndex < len(mm) && columIndex >= 0 {
		if numbers < mm[rowIndex][columIndex] {
			columIndex--
		} else if numbers > mm[rowIndex][columIndex] {
			rowIndex++
		} else {
			return true
		}
	}
	return false
}
