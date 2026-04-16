
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