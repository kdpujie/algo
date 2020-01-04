package linked

import (
	"fmt"
)

//存储的数据类型以int为例
//单链表节点
type Node struct {
	data byte
	next *Node
}

type LinkList struct {
	head *Node
}

//表尾添加元素
func (l *LinkList) add(value byte) {
	node := &Node{
		data: value,
	}
	if l.head == nil { //空链表
		l.head = node
	} else { //原链表非空
		index := l.head //从头开始遍历，找到尾节点
		for index.next != nil {
			index = index.next
		}
		index.next = node
	}
}

//遍历整个链表并打印
func (l *LinkList) traversal() {
	index := l.head
	for index != nil {
		fmt.Printf("%d,", index.data)
		index = index.next
	}
	fmt.Println()
}

//查找第一个值为value的节点,返回节点的地址,找不到返回NULL
func (l *LinkList) find(value byte) *Node {
	index := l.head
	for index != nil && index.data != value {
		index = index.next
	}
	return index
}

//删除第一个值为x的节点
func (l *LinkList) delete(value byte) {
	if l.head.data == value { //删除的是头节点
		l.head = l.head.next
	} else { //非头结点删除
		index := l.head
		for index.next != nil {
			if index.next.data == value { //找到第一个值等于value的节点后，断链删除该节点。
				index.next = index.next.next
				return
			}
			index = index.next
		}
	}
}

//在节点p后插入值为x的节点
func (l *LinkList) insert(p *Node, value byte) {
	node := &Node{
		data: value,
	}
	node.next = p.next
	p.next = node
	if l.head == nil {
		l.head = p
	}
}

//在链表的头部插入节点
func (l *LinkList) insertHead(value byte) {
	node := &Node{
		data: value,
	}
	node.next = l.head
	l.head = node
}

// 单链表反转
func (l *LinkList) reverse(head *Node) *Node {
	var pre, next *Node
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}

// 从头遍历链表
func (l *LinkList) print(node *Node) string {
	var str string
	index := node
	if nil == index {
		index = l.head
	}
	for index != nil {
		str = str + string(index.data)
		index = index.next
	}
	return str
}

//递归后续打印单链表
func (l *LinkList) reversePrint(node *Node) {
	if node == nil {
		return
	}
	if node.next == nil {
		fmt.Printf("%d,", node.data)
		return
	} else {
		l.reversePrint(node.next)
		fmt.Printf("%d,", node.data)
	}

}
