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

func construct(inorderStart int, inorderEnd int, preorder []int, inorderMap map[int]int, preorderIdx *int) *TreeNode {
	if inorderStart > inorderEnd {
		return nil
	}

	rootVal := preorder[*preorderIdx]
    root := &TreeNode{Val: rootVal}
    *preorderIdx++
	mid := inorderMap[rootVal]
	root.Left = construct(inorderStart, mid-1, preorder, inorderMap, preorderIdx)
	root.Right = construct(mid+1, inorderEnd, preorder, inorderMap, preorderIdx)
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i:=0; i<len(inorder); i++ {
		inorderMap[inorder[i]] = i
	}
    
	preorderIdx := 0
	return construct(0, len(preorder)-1, preorder, inorderMap, &preorderIdx)
}

func main() {
    preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}


    fmt.Println(buildTree(preorder, inorder))
}