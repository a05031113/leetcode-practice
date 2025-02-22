package main

import (
	"fmt"
	"strings"
)


type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func newNode(val int) *TreeNode {
    return &TreeNode{Val: val}
}

func serializeTree(root *TreeNode) string {
    var result strings.Builder
    var stack []*TreeNode
    stack = append(stack, root)
    
    for len(stack) > 0 {
        curr := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        result.WriteString(fmt.Sprintf("^%d", curr.Val))
        
        if curr.Right != nil {
            stack = append(stack, curr.Right)
        } else {
            result.WriteString("*")
        }
        if curr.Left != nil {
            stack = append(stack, curr.Left)
        } else {
            result.WriteString("#")
        }
    }
    return result.String()
}


func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    rootStr := serializeTree(root)
    subStr := serializeTree(subRoot)
    return strings.Contains(rootStr, subStr)
}

func main() {
    root := newNode(4)
    root.Left = newNode(5)

	subRoot := newNode(4)
    subRoot.Right = newNode(5)

    fmt.Println(isSubtree(root, subRoot))
}