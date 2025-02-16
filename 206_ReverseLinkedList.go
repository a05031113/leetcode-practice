package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
	current := head

	for current != nil {
		nextTemp := current.Next //  &ListNode{Val: 2}
        current.Next = prev // nil
        prev = current //  &ListNode{Val: 1}
        current = nextTemp //  &ListNode{Val: 2}
	}
	return prev
}

func main() {
	head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 3}
    head.Next.Next.Next = &ListNode{Val: 4}
    head.Next.Next.Next.Next = &ListNode{Val: 5}

	printList := func(node *ListNode) {
        for node != nil {
            fmt.Printf("%d -> ", node.Val)
            node = node.Next
        }
        fmt.Println("nil")
    }

	printList(head)
	printList(reverseList(head))
	printList(head)
	// fmt.Println(reverseList(head))
}