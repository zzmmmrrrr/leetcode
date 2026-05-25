/*
思路概括：
1. 先计算链表总长度 length，确定需要反转的组数。
2. 外层循环按 k 为步长遍历每组，内层循环进行组内 k 次指针反转。
3. 组间连接：用 lastGroupTail 记录上一组的尾节点（也是当前组反转前的头节点前一个节点），将当前组反转后的新头节点接到 lastGroupTail 后，新尾节点指向下一组起点。

复杂度：
- 时间：O(N)，每个节点遍历并反转一次。
- 空间：O(1)，仅使用有限个指针变量。
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	//第一步先计算出链表长度
	length := 0
	for i := head; i != nil; i = i.Next {
		length++
	}
	//第二步初始化新链表相关节点
	//lastGroupTail会指向每组未反转时的前一个节点，或者说上一组的反转后的尾节点
	// 也就是这一组反转后新头节点要连接在后面的位置
	dummy := &ListNode{Next: head} //dummyNode
	lastGroupTail := dummy
	var pre *ListNode
	cur := head

	//第三步开始遍历，分为组内反转链表和组间连接
	//i相当于每一组的第一个位置的角标
	//类比成切片的话，i从0开始，步长为k，最后一组是i到i+k-1
	//i用来遍历组数，j用来进行组内k次反转
	for i := 0; i+k-1 < length; i += k {
		for j := 0; j < k; j++ {
			//反转固定步骤：
			//保留nxt-cur.Next指向pre-移动pre-移动cur
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}

		//组间连接
		//oldHead是反转前这一组的头节点，反转后需要充当这一组尾节点
		//cur此时指向下一组起点位置，因此oldHead要指向cur
		//pre此时指向这一组反转前的尾节点，此时是新的头节点
		oldHead := lastGroupTail.Next
		oldHead.Next = cur
		lastGroupTail.Next = pre
		lastGroupTail = oldHead
		pre = nil
		//保持逻辑一致，每一组开始前pre都保持nil
		//不这样做也不会出错，
		// 因为每一组的第一个被反转的节点一开始指向一个未知的pre，
		// 后续又会连接到下一组开头
	}

	return dummy.Next
}

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
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	list := buildList(nums)
	newList := reverseKGroup(list, 2)
	for head := newList; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}
