package main

import "fmt"

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

/*
判断 tree2 是否是 tree1 的子结构
*/
func preOrder(tree1, tree2 *TreeNode) bool {
	// 题目通常规定：空树不是任意树的子结构
	if tree1 == nil || tree2 == nil {
		return false
	}

	// 当前节点匹配
	// 或左子树存在匹配
	// 或右子树存在匹配
	return isContained(tree1, tree2) ||
		preOrder(tree1.Left, tree2) ||
		preOrder(tree1.Right, tree2)
}

/*
判断从当前节点开始是否匹配
*/
func isContained(p, q *TreeNode) bool {
	// q 已匹配完成
	if q == nil {
		return true
	}

	// p 已空，但 q 还没匹配完
	if p == nil {
		return false
	}

	// 当前值相等
	// 且左右子树都匹配
	return p.Val == q.Val &&
		isContained(p.Left, q.Left) &&
		isContained(p.Right, q.Right)
}

func main() {

	/*
	   tree1

	       1
	      / \
	     2   3
	    / \
	   4   5
	*/

	tree1 := &TreeNode{Val: 1}
	tree1.Left = &TreeNode{Val: 2}
	tree1.Right = &TreeNode{Val: 3}
	tree1.Left.Left = &TreeNode{Val: 4}
	tree1.Left.Right = &TreeNode{Val: 5}

	/*
	   tree2

	       2
	      / \
	     4   5
	*/

	tree2 := &TreeNode{Val: 2}
	tree2.Left = &TreeNode{Val: 4}
	tree2.Right = &TreeNode{Val: 5}

	res := preOrder(tree1, tree2)

	fmt.Println(res)
}
