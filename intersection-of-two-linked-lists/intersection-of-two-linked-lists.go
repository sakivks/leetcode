package main

import "fmt"

func main() {
	listx := &ListNode{2, &ListNode{4, nil}}
	list1 := &ListNode{1, &ListNode{9, listx}}
	list2 := &ListNode{3, listx}

	fmt.Println(getIntersectionNode(list1, list2)) // shoud return ListNode with value 2
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	if p1 == nil || p2 == nil {
		return nil
	}
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
		if p1 == p2 {
			return p1
		}
		if p1 == nil {
			p1 = headB
		}
		if p2 == nil {
			p2 = headA
		}
	}
	return p1
}

// brute force
func getIntersectionNode_brute(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	list2 := headB

	for headA != nil {
		for headB != nil {
			if headA == headB {
				return headA
			}
			headB = headB.Next
		}
		headB = list2
		headA = headA.Next
	}
	return nil
}
