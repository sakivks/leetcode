package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	head := mergeTwoLists(l1, l2)
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	head := result

	for l1 != nil || l2 != nil {
		var next *ListNode

		if l1 == nil {
			next = l2
			l2 = l2.Next
		} else if l2 == nil {
			next = l1
			l1 = l1.Next
		} else if l1.Val < l2.Val {
			next = l1
			l1 = l1.Next
		} else {
			next = l2
			l2 = l2.Next
		}

		head.Next = next
		head = head.Next
	}

	return result.Next
}
