package main

import "fmt"

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

	neededNextResNodeList := false

	for curList1Node.Next != nil || curList2Node.Next != nil || sum > 9 || extra {
		if extra {
			sum++
		}
		extra = sum > 9

		tail := sum % 10
		curResListNode.Val = tail

		if curList1Node.Next != nil {
			curList1Node = curList1Node.Next
			neededNextResNodeList = true
		} else {
			curList1Node = &ListNode{
				Val:  0,
				Next: nil,
			}
			neededNextResNodeList = false
		}

		if curList2Node.Next != nil {
			curList2Node = curList2Node.Next
			neededNextResNodeList = true
		} else {
			curList2Node = &ListNode{
				Val:  0,
				Next: nil,
			}
			neededNextResNodeList = false
		}

		if neededNextResNodeList {
			curResListNode.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			curResListNode = curResListNode.Next
		}

		sum = curList1Node.Val + curList2Node.Val
	}

	return resList
}

func main() {
	l1 := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	l2 := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	res := addTwoNumbers(&l1, &l2)
	fmt.Println(res)
}
