package main

import "log"

func main() {
	var str = " I am a    student"
	strB := []byte(str)
	mirroSentence(strB)
	log.Printf("after mirro:%s\n", string(strB))
	strB1 := reverse3(strB)
	log.Printf("reverse3:%s, %s\n", string(strB1), string(strB))
}

/**
1.题目描述
	给出一个任意英文语句，将其句子反转。比如: “I am a student”-> “student a am I”
2.思路1:
	将整个句子反转，然后再将单词反转。
**/

func mirroSentence(strB []byte) {
	if len(strB) < 1 {
		return
	}
	reverse2(strB, 0, len(strB)-1) //反转整个句子
	var start int
	//反转单词
	for index := 0; index < len(strB); index++ {
		if strB[index] == 32 && start != index {
			reverse2(strB, start, index-1)
			start = index
		}
		if strB[start] == 32 && strB[index] != 32 {
			start = index
		}
	}
}

/**
@description 字符串反转
	思路：头尾指针start，end一起相向移动，通过中间变量的方式交换start、end移动过程中的值，直到start不小于end。
**/
func reverse1(strB []byte, start, end int) {
	var temp byte
	for start < end {
		temp = strB[start]
		strB[start] = strB[end]
		strB[end] = temp
		start++
		end--
	}
}

/**
@description 字符串反转
	思路：和reverse1方法的思路一致，不同在于变量交换的方式。该方法通过异或(^)操作交换start、end的值，而不引入中间变量
**/
func reverse2(strB []byte, start, end int) {
	for start < end {
		strB[start] ^= strB[end]
		strB[end] ^= strB[start]
		strB[start] ^= strB[end]
		start++
		end--
	}
}

/**
@description 字符串反转
	思路：递归结束的临界条件为，只剩下一个字符；否则就把除第一个字符
**/
func reverse3(strB []byte) []byte {
	if len(strB) == 1 {
		return strB
	}
	return append(reverse3(strB[1:]), strB[:1]...)
}
