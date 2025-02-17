package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func reorderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
        return head
    }
	slow, fast := head, head
    var prev *ListNode
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    prev.Next = nil

	var reverseList *ListNode
	currentLarge := slow
	for currentLarge != nil {
		nextTemp := currentLarge.Next //  &ListNode{Val: 2}
        currentLarge.Next = reverseList // nil
        reverseList = currentLarge //  &ListNode{Val: 1}
        currentLarge = nextTemp //  &ListNode{Val: 2}
	}

	first := head
    second := reverseList

	for first != nil && second != nil {
        firstNext := first.Next
        secondNext := second.Next
        
        first.Next = second
		if firstNext != nil {
			second.Next = firstNext
		}
        
        first = firstNext
        second = secondNext
    }

	return head
}


func main() {
	head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 3}
    head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	// head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	printList := func(node *ListNode) {
        for node != nil {
            fmt.Printf("%d -> ", node.Val)
            node = node.Next
        }
        fmt.Println("nil")
    }

	printList(reorderList(head))
	// fmt.Println(reverseList(head))
}