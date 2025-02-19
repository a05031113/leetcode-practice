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


func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
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
    
    root := newNode(1)
    root.Left = newNode(2)
    root.Right = newNode(3)
    root.Left.Left = newNode(4)
    root.Right.Right = newNode(5)
    root.Right.Right.Right = newNode(6)

    // Calculate and print the maximum depth
    depth := maxDepth(root)
    fmt.Printf("Maximum depth of the tree is: %d\n", depth)

    // Create and test another tree
    //     1
    //    /
    //   2
    root2 := newNode(1)
    root2.Left = newNode(2)
    
    depth2 := maxDepth(root2)
    fmt.Printf("Maximum depth of the second tree is: %d\n", depth2)
}