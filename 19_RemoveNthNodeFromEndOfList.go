package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    fast, slow := head, head

	for i:=0; i<n; i++{
		fast = fast.Next
	}

	if fast == nil && n == 1{
		return nil
	} else if fast == nil {
        return slow.Next
    }

	for fast != nil {
		fast = fast.Next

		if fast == nil {
			slow.Next = slow.Next.Next
			break
		}
		slow = slow.Next
	} 

	return head
}

func main() {
	head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 3}
    head.Next.Next.Next = &ListNode{Val: 4}
    head.Next.Next.Next.Next = &ListNode{Val: 5}

	n:=2

	printList := func(node *ListNode) {
        for node != nil {
            fmt.Printf("%d -> ", node.Val)
            node = node.Next
        }
        fmt.Println("nil")
    }

	printList(removeNthFromEnd(head, n))
	// fmt.Println(reverseList(head))
}