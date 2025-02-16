package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    output := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			output.Next = list1
			list1 = list1.Next
		} else {
			output.Next = list2
			list2 = list2.Next
		}
		output = output.Next
	}

	if list1 != nil {
        output.Next = list1
    }
    if list2 != nil {
        output.Next = list2
    }
	return dummy.Next
}

func main() {
	list1 := &ListNode{Val: 1}
    list1.Next = &ListNode{Val: 2}
    list1.Next.Next = &ListNode{Val: 4}

	list2 := &ListNode{Val: 1}
    list2.Next = &ListNode{Val: 3}
    list2.Next.Next = &ListNode{Val: 4}

	printList := func(node *ListNode) {
        for node != nil {
            fmt.Printf("%d -> ", node.Val)
            node = node.Next
        }
        fmt.Println("nil")
    }

	printList(mergeTwoLists(list1, list2))
	// fmt.Println(reverseList(head))
}