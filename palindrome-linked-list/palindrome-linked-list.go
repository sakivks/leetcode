package main

import "fmt"

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{1, nil}}}
	fmt.Println(isPalindrome(list))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// use two pointers to locate the mid-point of linked list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// skip central node if length of linked list is odd number
	if fast != nil {
		slow = slow.Next
	}

	tail := reverseList(slow)

	for tail != nil {
		if tail.Val != head.Val {
			return false
		}
		head, tail = head.Next, tail.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

//Solution with O(n) space
func isPalindrome_spaceOn(head *ListNode) bool {
	reversedList := reverseList_copy(head)

	for head != nil {
		fmt.Println(head.Val, reversedList.Val)
		if head.Val != reversedList.Val {
			return false
		}
		head, reversedList = head.Next, reversedList.Next
	}

	return true
}

func reverseList_copy(head *ListNode) *ListNode {
	var prev *ListNode
	cur := &ListNode{head.Val, head.Next}
	for cur != nil {
		if cur.Next != nil {
			cur.Next, prev, cur = prev, cur, &ListNode{cur.Next.Val, cur.Next.Next}
		} else {
			cur.Next, prev, cur = prev, cur, nil
		}
	}
	return prev
}
