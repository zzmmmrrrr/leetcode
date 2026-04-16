package main

import (
	"fmt"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	//在计算高度的同时计算出ans
	//相当于每一个节点都有可能是直径所在的子树根节点
	//在遍历高度的同时找出能使路径最大的拐点，此时的路径就是直径
	//这里取根节点的高度为1
	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftH := height(node.Left)
		rightH := height(node.Right)
		ans = max(ans, leftH+rightH)
		return max(leftH, rightH) + 1
	}
	height(root)
	return ans
}

func main() {
	head := &TreeNode{Val: 1}
	head.Left = &TreeNode{Val: 2}
	head.Right = &TreeNode{Val: 3}
	x := head.Left
	x.Left = &TreeNode{Val: 4}
	x.Right = &TreeNode{Val: 5}
	ans := diameterOfBinaryTree(head)
	fmt.Println(ans)
}
