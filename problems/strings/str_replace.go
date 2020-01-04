package main

import "log"

func main() {
	var str = "I am a Student."
	newStrB := strReplace2([]byte(str))
	log.Printf("newStr=%s\n", string(newStrB))
}

/**
1.问题描述:
	实现一个函数，把字符串中的每个空格替换成"%20"。例如输入"We are happy."，则输出"We%20are%20happy."
	空格的ASCII码为32
	%的ASCII码为37
2.思路1：
	空格占一个字符，"%20"占三个字符，遍历字符串，遇到空格后把后续字符后移三个位置，空出来的位置补"%20"。时间复杂度为O(n^2)
3.思路2：
	先遍历一遍数组，计算出空格数量。然后扩展数组
	设置两个指针P1，P2， P1指向扩展前的末尾，用于寻找空格；P2指向扩展后的末尾，用于在P1发现空格后，自身位置向前移动3个字符，三个字符为"%20"。遍历完整后，P1和P2在结尾处汇合。
	时间复杂度O(2n)

**/
func strReplace2(strB []byte) []byte {
	var blankNum, originalLength int
	for _, v := range strB {
		originalLength++
		if v == 32 {
			blankNum++
			strB = append(strB, 32, 32) //扩容
		}
	}
	if blankNum == 0 {
		return strB
	}
	originalLastIndex := originalLength - 1
	newLastIndex := originalLastIndex + blankNum*2
	for ; originalLastIndex >= 0; originalLastIndex-- {
		if strB[originalLastIndex] == 32 {
			strB[newLastIndex] = 48
			newLastIndex--
			strB[newLastIndex] = 50
			newLastIndex--
			strB[newLastIndex] = 37
		} else {
			strB[newLastIndex] = strB[originalLastIndex]
		}
		newLastIndex--
	}
	return strB
}
