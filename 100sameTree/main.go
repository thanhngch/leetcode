package main

import "fmt"

// TreeNode is tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		return p.Val == q.Val && isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right)
	}
	return false
}

func main() {
	t1 := &TreeNode{Val: 1}
	t1.Left = &TreeNode{Val: 2}
	t1.Right = &TreeNode{Val: 3}

	t2 := &TreeNode{Val: 1}
	t2.Left = &TreeNode{Val: 2}
	t2.Right = &TreeNode{Val: 3}
	fmt.Println(isSameTree(t1, t2))
}
