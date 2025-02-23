package main

import (
	"fmt"
)


type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func newNode(val int) *TreeNode {
    return &TreeNode{Val: val}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}



func main() {
    root := newNode(6)
    root.Left = newNode(2)
	root.Right = newNode(8)
	root.Left.Left = newNode(0)
	root.Left.Right = newNode(4)
	root.Left.Right.Left = newNode(3)
	root.Left.Right.Right = newNode(5)
	root.Right.Left = newNode(7)
	root.Right.Right = newNode(9)

	p := newNode(2)
	q := newNode(4)

    fmt.Println(lowestCommonAncestor(root, p, q))
}