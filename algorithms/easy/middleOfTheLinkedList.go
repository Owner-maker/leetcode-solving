package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	curNode := head

	if curNode.Next == nil {
		return head
	}

	amount := 1
	for curNode.Next != nil {
		curNode = curNode.Next
		amount++
	}

	middle := amount/2 + 1
	curNode = head
	amount = 1

	for curNode.Next != nil {
		if amount == middle {
			return curNode
		}
		amount++
		curNode = curNode.Next
	}
	if amount == middle {
		return curNode
	}

	return nil
}
