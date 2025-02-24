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

func levelOrder(root *TreeNode) [][]int {
	var output [][]int
    if root == nil {
        return output
    }
	var stack []map[*TreeNode]int
	m := make(map[*TreeNode]int)
	m[root] = 0
    stack = append(stack, m)
	levelMap := make(map[int][]int)

	for len(stack) > 0 {
		currMap := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for node, level := range currMap {
			if _, exist := levelMap[level]; exist {
				levelMap[level] = append(levelMap[level], node.Val)
			} else {
				levelMap[level] = []int{node.Val}
			}
			if node.Right != nil {
				m := make(map[*TreeNode]int)
				m[node.Right] = level + 1
				stack = append(stack, m)
			}
			if node.Left != nil {
				m := make(map[*TreeNode]int)
				m[node.Left] = level + 1
				stack = append(stack, m)
			}
		}
	}
	for i:=0; i<len(levelMap); i++ {
		output = append(output, levelMap[i])
	}

	return output
}



func main() {
    root := newNode(3)
    root.Left = newNode(9)
	root.Right = newNode(20)
	root.Right.Left = newNode(15)
	root.Right.Right = newNode(7)

    fmt.Println(levelOrder(root))
}