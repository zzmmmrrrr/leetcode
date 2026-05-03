/*
思路概括：
1. 递归终止条件：节点为空或单个节点时直接返回自身。
2. 递归反转 head.Next 之后的子链表，得到子链表的新头节点 newHead。
3. 将 head 接到其后继之后（head.Next.Next = head），并断开原来的链（head.Next = nil）。
4. newHead 从最后一层递归向上传递，最终即为全局新头节点。

复杂度：
- 时间：O(N)，每个节点访问一次。
- 空间：O(N)，递归调用栈深度等于链表长度。
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

//第一步：
//确定函数功能
//接收一个链表的头节点，返回反转后链表的新头节点

//第二步：
//默认函数已经黑盒实现，思考当前节点和子问题的关系
//比如当前节点cur和nxt的关系
//如果把nxt交给函数，那么会返回nxt后局部链表的反转版本头节点
//此时cur应该接到nxt后面

// 第三步：
// 确定边界条件
// 链表只剩0个或者1个节点时无需反转，直接返回自身即可
// 也即node==nil或者node.Next==nil时返回node
func reverse(node *ListNode) *ListNode {
	//函数里面的书写流程是321

	if node == nil || node.Next == nil {
		return node
	}

	newHead := reverse(node.Next)
	node.Next.Next = node
	node.Next = nil

	//这里的newHead就是最终反转后全局新头节点
	//原因是递归不断深入到最后一个节点，返回最后一个节点作为newHead
	//每一层函数都是传递同一个newHead到上一层调用
	return newHead
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
