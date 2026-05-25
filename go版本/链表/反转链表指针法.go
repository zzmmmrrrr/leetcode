/*
思路概括：
1. 用三个指针 pre（已反转部分的头）、cur（当前待反转节点）、nxt（暂存 cur 的下一个节点）。
2. 每次迭代：保存 nxt = cur.Next，将 cur.Next 指向 pre，pre 后移，cur 后移。
3. 当 cur 为 nil 时结束，pre 指向原链表的尾节点，即新链表的头节点。

复杂度：
- 时间：O(N)，遍历一次链表。
- 空间：O(1)，仅使用有限个指针。
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
var pre *ListNode
pre ──→ nil（什么都没有）

pre := &ListNode{}
pre ──→ [ Val:0 | Next:nil ]（真实存在的节点）
*/
func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	//等价于pre:=(*ListNode)(nil)
	//这里不能把pre初始化为&ListNode{}，这样做会导致pre节点的Val==0
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// 创建链表时利用dummy哑节点，避免头节点特殊处理
// 原因是头节点为nil时不能做.Next操作获取字段
func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, x := range nums {
		cur.Next = &ListNode{Val: x}
		cur = cur.Next
	}
	return dummy.Next
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := buildList(nums)
	newHead := reverse(head)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}
