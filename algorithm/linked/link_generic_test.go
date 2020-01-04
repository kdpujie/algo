package linked

import "testing"

func TestAdd(t *testing.T) {
	var list LinkList
	list.add(1)
	list.add(2)
}

func TestReverse(t *testing.T) {
	var list LinkList
	list.add('a')
	list.add('b')
	list.add('c')
	h := list.reverse(list.head)
	if list.print(h) != "cba" {
		t.Error("reverse() failed. Got", list.print(h), "Expected cba")
	}
}

func TestIsPalindrome(t *testing.T) {
	var list LinkList
	list.add('a')
	list.add('b')
	list.add('c')
	list.add('b')
	list.add('a')
	res := isPalindrome(list)
	t.Logf("%s 是否是回文：%v \n", list.print(nil), res)
	if !res {
		t.Error("IsPalindrome() failed.")
	}
}
