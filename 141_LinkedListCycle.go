package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	head := &ListNode{Val: 3}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 0}
    head.Next.Next.Next = &ListNode{Val: -4}

	printList := func(node *ListNode) {
        for node != nil {
            fmt.Printf("%d -> ", node.Val)
            node = node.Next
        }
        fmt.Println("nil")
    }

	printList(reverseList(head))
	// fmt.Println(reverseList(head))
}