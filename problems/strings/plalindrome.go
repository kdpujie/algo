package main

import "log"

func main() {
	var str = "ABCDBA"
	plalindrome(str)
}

/**
1.题目描述
	给出一个长度不超过1000的字符串，判断它是不是回文(顺读，逆读均相同)的。
	可能有多组测试数据，对于每组数据，如果是回文字符串则输出"Yes!”，否则输出"No!"。
2.思路:
	遍历输入的字符串，设置两个指针start、end分别指向第一个字符和最后一个字符，循环结束条件为start < end ；如果start位置的字符!=end位置的内容，则提前结束循环，输出No！
	确定清楚，空串是否是回文。这里代码实现为空串是回文。
**/
func plalindrome(str string) {
	for i := 0; i < len(str)-1-i; i++ {
		if str[i] != str[len(str)-1-i] {
			log.Printf("%s是否是回文：%v\n", str, false)
			return
		}
	}
	log.Printf("%s是否是回文：%v\n", str, true)
}
