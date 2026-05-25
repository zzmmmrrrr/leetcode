package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1-2-3-4-5-6-7-nil
// k=3,要找倒数第三个节点,也即val==5
// 可以先让fast走k步，再让fast和slow一起走，直到fast遇到nil
// 相当于总共移动了n+n-k步
func solve(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	dummy := &ListNode{}
	cur := dummy
	for _, x := range nums {
		cur.Next = &ListNode{Val: x}
		cur = cur.Next
	}
	head := dummy.Next
	k := 3
	node := solve(head, k)
	fmt.Println(node.Val)
}
