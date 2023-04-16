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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	resList := &ListNode{
		Val:  0,
		Next: nil,
	}

	curList1Node := l1
	curList2Node := l2
	curResListNode := resList

	sum := l1.Val + l2.Val
	extra := false

	for curList1Node.Next != nil || curList2Node.Next != nil || sum > 0 || extra {
		if extra {
			sum++
		}
		extra = sum > 9

		curResListNode.Val = sum % 10
		sum = 0
		if curList1Node.Next != nil || curList2Node.Next != nil || sum > 9 || extra {
			curResListNode.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			curResListNode = curResListNode.Next
		}

		if curList1Node.Next != nil {
			curList1Node = curList1Node.Next
			sum += curList1Node.Val
		}

		if curList2Node.Next != nil {
			curList2Node = curList2Node.Next
			sum += curList2Node.Val
		}
	}
	curResListNode = nil

	return resList
}
