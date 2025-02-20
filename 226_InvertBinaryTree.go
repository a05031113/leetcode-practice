package main

import "fmt"


type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func newNode(val int) *TreeNode {
    return &TreeNode{Val: val}
}


func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}
	left := root.Left
	right := root.Right
	root.Left = right
	root.Right = left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}


func main() {
    // Create a sample binary tree:
    //       1
    //      / \
    //     2   3
    //    /     \
    //   4       5
    //            \
    //             6
    
    root := newNode(4)
    root.Left = newNode(2)
    root.Right = newNode(7)
    root.Left.Left = newNode(1)
	root.Left.Right = newNode(3)
    root.Right.Left = newNode(6)
    root.Right.Right = newNode(9)

    
    fmt.Println(invertTree(root))
}