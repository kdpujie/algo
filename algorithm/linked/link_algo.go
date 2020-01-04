package linked

/**
@description
	用单链表的结构存储数据，判断输入字符串是否为回文
时间复杂度：O(n)
空间复杂度：O(1)
实现思路：
	1. 快慢指针找到中点
	2. 后半段逆序和前半段比较
	3. 后半段恢复原状
**/
func isPalindrome(h LinkList) bool {
	var slow, fast, head = h.head, h.head, h.head
	for nil != fast.next && nil != fast.next.next {
		slow = slow.next
		fast = fast.next.next
	}
	h1 := h.reverse(slow.next)
	res := true
	h2 := h1
	for nil != h1 && nil != head {
		if head.data != h1.data {
			res = false
			break
		}
		h1 = h1.next
		head = head.next
	}
	slow.next = h.reverse(h2)
	return res
}
