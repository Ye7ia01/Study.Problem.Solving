package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	ans := &ListNode{} // ans type *ListNode
	tail := ans        // copy address of ans to tail
	for list1 != nil && list2 != nil {

		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next

		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// Append the remaining elements of list1 or list2, if any
	for list1 != nil {
		tail.Next = list1
		tail = tail.Next
		list1 = list1.Next
	}

	for list2 != nil {
		tail.Next = list2
		tail = tail.Next
		list2 = list2.Next
	}

	return ans.Next // head of the merged linked list
}

func main() {

	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	list2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}
	merged := mergeTwoLists(list1, list2)
	for merged != nil {
		fmt.Println(merged.Val)
		merged = merged.Next
	}
}
