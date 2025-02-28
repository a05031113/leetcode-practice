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

// func bst(root *TreeNode, nums *[]int) {
// 	if root == nil {
// 		return 
// 	}
// 	bst(root.Left, nums)
// 	*nums = append(*nums, root.Val)
// 	bst(root.Right, nums)
// }

// func kthSmallest(root *TreeNode, k int) int {
//     var nums []int
// 	bst(root, &nums)
// 	return nums[k-1]
// }

func kthSmallest(root *TreeNode, k int) int {
    // Stack to simulate recursion
    var stack []*TreeNode
    curr := root
    count := 0
    
    // Inorder traversal using iteration
    for curr != nil || len(stack) > 0 {
        // Go left as far as possible
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        
        // Pop from stack
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        // Count this node
        count++
        if count == k {
            return curr.Val
        }
        
        // Go right
        curr = curr.Right
    }
    
    return -1 // If k is larger than the tree size
}


func main() {
    root := newNode(5)
    root.Left = newNode(3)
	root.Right = newNode(6)
	root.Left.Left = newNode(2)
	root.Left.Right = newNode(4)
	root.Left.Left.Left = newNode(1)


    fmt.Println(kthSmallest(root, 3))
}