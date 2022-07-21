package main

//返回链表的逆序结构
//可以使用迭代和递归两种，你是否都可以实现？
type ListNode struct {
	Val  int
	Next *ListNode
}

//a -> b -> c -> d
//
//a <- b -> c -> d
//
//
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next // 工具人变量，完全是为了移动cur用的
		cur.Next = pre   // 关键步骤：修改指针方向
		pre = cur
		cur = next
	}
	return pre
}
