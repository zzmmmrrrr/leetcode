
/*
思路概括：
1. 对 tree1 进行先序遍历，在每个节点处调用 isContained 判断以该节点为根的树是否包含 tree2 的结构。
2. isContained 递归判断：若 tree2 为 nil 则包含成立；若 tree1 为 nil 而 tree2 非 nil 则失败；否则需值相等且左右子树分别包含。

复杂度：
- 时间：O(M × N)，M 为 tree1 节点数，N 为 tree2 节点数，最坏每个 tree1 节点都需调用 isContained。
- 空间：O(max(H1, H2))，递归栈深度。
*/

func preOrder(tree1,tree2 *TreeNode){
	if tree1==nil || tree2==nil{
		return false
	}
	return isContained(tree1,tree2) || preOrder(tree1.Left,tree2)||preOrder(tree1.Right,tree2)

}

func isContained(p,q *TreeNode) bool{
	if q==nil{
		return true
	}
	if p==nil{
		return false
	}
	return p.Val==q.Val && isContained(p.Left,q.Left) &&isContained(p.Right,q.Right)
}