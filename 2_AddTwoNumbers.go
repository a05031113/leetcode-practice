package main

import (
    "fmt"
    "math"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    carry := 0
    for l1 != nil || l2 != nil || carry > 0 {
        num1 := 0
        if l1 != nil {
            num1 = l1.Val
            l1 = l1.Next
        }

        num2 := 0
        if l2 != nil {
            num2 = l2.Val
            l2 = l2.Next
        }

        sum := num1 + num2 + carry
        carry = sum / 10
        digit := sum % 10

        current.Next = &ListNode{Val: digit}
        current = current.Next
    }
    return dummy.Next
}

func printList(head *ListNode) {
    current := head
    for current != nil {
        fmt.Printf("%d -> ", current.Val)
        current = current.Next
    }
    fmt.Println("nil")
}

func main() {
    // Create test linked lists
    // Example: 2 -> 4 -> 3 (representing 342)
    l1 := &ListNode{Val: 2}
    l1.Next = &ListNode{Val: 4}
    l1.Next.Next = &ListNode{Val: 3}

    // Example: 5 -> 6 -> 4 (representing 465)
    l2 := &ListNode{Val: 5}
    l2.Next = &ListNode{Val: 6}
    l2.Next.Next = &ListNode{Val: 4}

    // fmt.Println("First number:")
    // printList(l1)
    // fmt.Println("Second number:")
    // printList(l2)

    result := addTwoNumbers(l1, l2)
    fmt.Println("Result:")
    printList(result)
}