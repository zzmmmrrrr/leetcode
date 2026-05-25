/*
思路概括：
1. 前序遍历（preOrder）：先访问根节点，再递归左子树，最后递归右子树。
2. 中序遍历（inorderTraversal）：先递归左子树，再访问根节点，最后递归右子树。
3. 使用闭包递归函数收集节点值到结果切片中。

复杂度：
- 时间：O(N)，每个节点访问一次。
- 空间：O(H)，H 为树高，递归栈深度（不计结果数组）。
*/
package main

func preOrder(root *TreeNode) []int {
	res:=make([]int,0)
	var dfs (root *TreeNode) 
	dfs=func(root *TreeNode){
		if root==nil{
			return
		}
		res=append(res,root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res:=make([]int,0,110)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res


}

