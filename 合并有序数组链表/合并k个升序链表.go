/*
思路概括（两种方法）：
1. 小顶堆法：将每个链表的头节点加入最小堆，每次弹出堆顶作为结果的下一个节点，并将该节点的 Next 入堆，直到堆为空。
2. 分治法：用二分归并将 k 个链表递归拆分，底层利用合并两个有序链表的基础操作。

复杂度：
- 小顶堆法：时间 O(N log K)，空间 O(K)，N 为所有节点总数，K 为链表数。
- 分治法：时间 O(N log K)，空间 O(log K)（递归栈深度）。
*/

package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type hp []*ListNode

//heap要实现五个方法，Less,Len,Swap,Pop,Push
//其中要修改hp长度的必须要传入*h，并手动解指针

// Push接收泛型，压入堆时要断言
// Pop弹出泛型，使用时要断言
func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(*ListNode))
}

func (h *hp) Pop() any {
	a := *h
	n := len(a)
	top := a[n-1]
	*h = a[:n-1]
	return top
}
func mergeKLists(lists []*ListNode) *ListNode {
	//1.把每个链表的首个元素加入堆中,并初始化堆
	h := make(hp, 0, len(lists))

	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}
	heap.Init(&h)

	//2.利用dummynode，创建一个答案链表，用尾接法插入对应节点
	dummy := &ListNode{}
	cur := dummy

	//3.循环弹出堆顶，并把出堆的列表中下一个元素加入堆中
	for len(h) > 0 {
		top := heap.Pop(&h).(*ListNode)
		cur.Next = top
		cur = cur.Next
		if top.Next != nil {
			heap.Push(&h, top.Next)
		}
	}
	return dummy.Next

}

// ----------------------------------------------------------------------------------------------------------
// 分治法：用索引代替切片，避免重复遍历时链表被破坏
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var solve func(int, int) *ListNode //solve同时完成拆分和合并
	//left,right是lists列表里面的idx，表明left和right 位置之间的链表要被合并
	solve = func(left, right int) *ListNode {
		if left == right {
			return lists[left]
		}

		mid := left + (right-left)/2

		l := solve(left, mid)
		r := solve(mid+1, right)

		return mergeTwoLists(l, r)
	}

	return solve(0, len(lists)-1)
}

// 合并两个有序链表
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next
}

func main() {
	n1 := []int{1, 2, 3, 4}
	n2 := []int{0, 4, 5, 8}
	n3 := []int{6, 7, 9, 10000}
	lists := buildLists(n1, n2, n3)
	lists1 := buildLists(n1, n2, n3)
	merged := mergeKLists(lists)
	merged1 := mergeKLists1(lists1)
	fmt.Println("下面是小顶堆方法")
	for merged != nil {
		fmt.Println(merged.Val)
		merged = merged.Next
	}
	fmt.Println("下面是分治法")
	for merged1 != nil {
		fmt.Println(merged1.Val)
		merged1 = merged1.Next
	}

}

// 将多个 []int 转换为 []*ListNode（每个 []int 构成一个链表）
func buildLists(slices ...[]int) []*ListNode {
	lists := make([]*ListNode, 0, len(slices))
	for _, slice := range slices {
		dummy := &ListNode{}
		cur := dummy
		for _, val := range slice {
			cur.Next = &ListNode{Val: val}
			cur = cur.Next
		}
		lists = append(lists, dummy.Next) // dummy.Next 是链表头（可能为 nil）
	}
	return lists
}
