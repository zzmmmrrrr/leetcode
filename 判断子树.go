package main
import (
	"fmt"	
)
type TreeNode struct{
	Val int
	Left,Right *TreeNode
}

func getHeight(root *TreeNode) int{
	if root==nil{
		return 0
	}
	leftH:=getHeight(root.Left)
	rightH:=getHeight(root.Right)
	return max(leftH,rightH)+1
}

func isSameTree(p,q *TreeNode) bool{
	if p==nil||q==nil{
		return p==q
	}

	return p.Val==q.Val && isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
}

func isSubtree(root,subRoot *TreeNode) bool{
	subH:=getHeight(subRoot)
	var dfs func(*TreeNode) (int,bool)

	//返回node所在高度和是否找到了能和子树匹配的部分
	dfs=func(node *TreeNode) (int,bool){
		if node==nil{
			return 0,false
		}
		//递归左右子树，拿到左右子树的高度和匹配结果
		leftH,leftFound:=dfs(node.Left)
		rightH,rightFound:=dfs(node.Right)
		//如果子树已经找到匹配了，直接短路返回
		if leftFound||rightFound{
			return 0,true
		}
		//计算当前节点的高度，匹配时才比较
		nodeH:=max(leftH,rightH)+1
		judge:=(nodeH==subH&&isSameTree(node,subRoot))
		return nodeH,judge
	}	

	_,found:=dfs(root)
	return found
}