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


func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p != nil && q == nil {
        return false
    } else if p == nil && q != nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
	var pstack []*TreeNode
	var qstack []*TreeNode
	pstack = append(pstack, p)
	qstack = append(qstack, q)

    for len(pstack) > 0 {
		pcurr := pstack[len(pstack)-1]
		pstack = pstack[:len(pstack)-1]

		qcurr := qstack[len(qstack)-1]
		qstack = qstack[:len(qstack)-1]

		if pcurr.Val != qcurr.Val {
			return false
		}
        if pcurr.Right != nil && qcurr.Right != nil {
            pstack = append(pstack, pcurr.Right)
			qstack = append(qstack, qcurr.Right)
        } else if pcurr.Right != nil && qcurr.Right == nil {
			return false
		} else if pcurr.Right == nil && qcurr.Right != nil {
			return false
		}
		if pcurr.Left != nil && qcurr.Left != nil {
			pstack = append(pstack, pcurr.Left)
			qstack = append(qstack, qcurr.Left)
		} else if pcurr.Left != nil && qcurr.Left == nil {
			return false
		} else if pcurr.Left == nil && qcurr.Left != nil {
			return false
		}
    }
	return true
}

func main() {
    p := newNode(1)
    p.Left = newNode(2)
    p.Right = newNode(3)

	q := newNode(1)
    q.Left = newNode(2)
    q.Right = newNode(3)

    fmt.Println(isSameTree(p, q))
}