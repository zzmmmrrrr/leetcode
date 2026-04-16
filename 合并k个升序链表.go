package main

import (
	"container/heap"
	"fmt"
)

//要实现heap接口，需要实现len,swap,less,push,pop五个函数
//易错点是push和pop要实现泛型以及断言
//函数声明时，push必须接收泛型，pop必须返回泛型
//pop出元素后必须类型断言
type ListNode struct {
	Val  int
	Next *ListNode
}
type hp []*ListNode

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(*ListNode))
}

func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func mergeKLists(lists []*ListNode) *ListNode {
	//1.把k个链表的第一个元素加入到堆中
	h:=make(hp,0,len(lists))
	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}

	//2.
	heap.Init(&h)
	dummy := &ListNode{}
	cur := dummy
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
func main() {
	n1 := []int{1, 2, 3, 4}
	n2 := []int{0, 4, 5, 8}
	n3 := []int{6, 7, 9, 10000}
	n4 := []int{23, 121, 332}
	n5 := []int{1, 2, 3, 4}
	n6 := []int{-1, 0, 2, 444545}
	lists := buildLists(n1, n2, n3)
	lists1 := buildLists(n4, n5, n6)
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

// ----------------------------------------------------------------------------------------------------------
// 分治法：用索引代替切片，避免重复遍历时链表被破坏
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var merge func(int, int) *ListNode
	merge = func(left, right int) *ListNode {
		if left == right {
			return lists[left]
		}

		mid := left + (right-left)/2

		l := merge(left, mid)
		r := merge(mid+1, right)

		return mergeTwoLists(l, r)
	}

	return merge(0, len(lists)-1)
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
