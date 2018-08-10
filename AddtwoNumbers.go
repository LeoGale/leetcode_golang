package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 = &ListNode{0, nil}
	var l3Cur = l3
	var offset int
	for l1 != nil || l2 != nil {
		var Val = offset
		if l1 != nil {
			Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			Val += l2.Val
			l2 = l2.Next
		}

		l3Cur.Val = Val % 10
    
		if Val >= 10 {
			offset = 1
		} else {
      offset = 0
    }

		if l1 != nil || l2 != nil {
			l3Cur.Next = &ListNode{0, nil}
			l3Cur = l3Cur.Next
		} else if offset == 1 {
			l3Cur.Next = &ListNode{1, nil}
			l3Cur = l3Cur.Next
		}
	}
	return l3
}

func main() {
	var l1 = &ListNode{9, nil}
	var l12 = &ListNode{4, l1}
	var l123 = &ListNode{3, l12}

	var l2 = &ListNode{5, nil}
	var l22 = &ListNode{6, l2}
	var l223 = &ListNode{4, l22}
	var l3 = addTwoNumbers(l123, l223)

	for l3 != nil {
		fmt.Printf("%d ", l3.Val)
		l3 = l3.Next
	}
	var l4 = &ListNode{5, l1}
	var l5 = &ListNode{5, nil}
	var l6 = addTwoNumbers(l4, l5)

	for l6 != nil {
		fmt.Printf("%d ", l6.Val)
		l6 = l6.Next
	}
}
