package main

import "fmt"

type Node struct {
	Left, Right *Node
	Val         int
}

func main() {
	list := []int{1, 2, 3, 4, -1, -1, 5, 6}
	root := buildTree(list)
	ans := levelOrder(root)
	fmt.Println(ans)
}

func buildTree(list []int) *Node { //-1表示空节点
	if len(list) == 0 || list[0] == -1 {
		return nil
	}
	root := &Node{Val: list[0]}
	queue := []*Node{root}

	i := 1
	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]

		if i < len(list) {
			if list[i] != -1 {
				parent.Left = &Node{Val: list[i]}
				queue = append(queue, parent.Left)
			}
			i++
		}

		if i < len(list) {
			if list[i] != -1 {
				parent.Right = &Node{Val: list[i]}
				queue = append(queue, parent.Right)
			}
			i++
		}

	}

	return root
}
func levelOrder(root *Node) []int {
	queue := []*Node{root}
	ans := []int{}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:len(queue)]
		ans = append(ans, top.Val)
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return ans
}
