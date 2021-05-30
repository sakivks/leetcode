package main

import "fmt"

func main() {
	// list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	list := &ListNode{1, nil}

	fmt.Println(reverseList(list))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// elegant
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

// more detailed
func reverseList_detailed(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, next *ListNode
	cur := head

	for cur.Next != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	cur.Next = prev
	return cur
}
