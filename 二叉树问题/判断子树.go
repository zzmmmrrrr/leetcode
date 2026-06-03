package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/*
计算树高度
*/
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftH := getHeight(root.Left)
	rightH := getHeight(root.Right)

	return max(leftH, rightH) + 1
}

/*
判断两棵树是否完全相同
*/
func isSameTree(p, q *TreeNode) bool {

	// 两个都为空
	if p == nil && q == nil {
		return true
	}

	// 一个为空，一个不为空
	if p == nil || q == nil {
		return false
	}

	// 当前节点值相同
	// 且左右子树完全相同
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

/*
判断 subRoot 是否是 root 的子树
*/
func isSubtree(root, subRoot *TreeNode) bool {

	// subRoot 的固定高度
	subH := getHeight(subRoot)

	// DFS 搜索
	var dfs func(*TreeNode) bool

	dfs = func(node *TreeNode) bool {

		// 搜索结束
		if node == nil {
			return false
		}

		// 当前节点高度
		nodeH := getHeight(node)

		// 高度相同才比较结构
		if nodeH == subH &&isSameTree(node, subRoot) {
			return true
		}

		// 去左右子树继续搜索
		return dfs(node.Left) || dfs(node.Right)
	}

	return dfs(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	/*
	   root

	       3
	      / \
	     4   5
	    / \
	   1   2
	*/

	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}

	/*
	   subRoot

	          4
	         / \
	        1   2
	*/

	subRoot := &TreeNode{Val: 4}
	subRoot.Left = &TreeNode{Val: 1}
	subRoot.Right = &TreeNode{Val: 2}

	fmt.Println(isSubtree(root, subRoot))
}
