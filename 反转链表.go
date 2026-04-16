package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1, 2, 3, 4}
	dummy := &ListNode{}
	cur := dummy
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	newHead := reverseList1(dummy.Next)
	for cur := newHead; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}

}

// func reverseList(head *ListNode) *ListNode {
// 	cur := head
// 	var pre *ListNode
// 	for cur != nil {
// 		nxt := cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = nxt
// 	}
// 	return pre
// }
func reverseList1(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	newNode := reverseList1(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newNode
}
