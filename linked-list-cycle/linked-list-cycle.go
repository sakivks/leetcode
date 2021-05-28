package main

import "fmt"

func main() {
	list := &ListNode{3, &ListNode{2, &ListNode{0, &ListNode{-4, nil}}}}
	list.Next.Next.Next.Next = list
	// var list *ListNode

	fmt.Println(hasCycle(list)) //should print true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	p1, p2 := head, head

	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			return true
		}
	}

	return false
}
