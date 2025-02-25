package main

import (
	"fmt"
	"math"
)


type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func newNode(val int) *TreeNode {
    return &TreeNode{Val: val}
}

func dfs(root *TreeNode, left int, right int) bool {
	if root == nil {
		return true
	}
	if root.Val <= left || root.Val >= right {
		return false
	}
	return dfs(root.Left, left, root.Val) && dfs(root.Right, root.Val, right)
}

func isValidBST(root *TreeNode) bool {
    return dfs(root, math.MinInt64, math.MaxInt64)
}

func main() {
    root := newNode(5)
    root.Left = newNode(4)
	root.Right = newNode(6)
	root.Right.Left = newNode(3)
	root.Right.Right = newNode(7)


    fmt.Println(isValidBST(root))
}