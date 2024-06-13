package p2

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
	carry := false
	val := 0

	fake := &ListNode{}
	next := fake
	for l1 != nil && l2 != nil {
		if carry {
			val = l1.Val + l2.Val + 1
		} else {
			val = l1.Val + l2.Val
		}
		if val >= 10 {
			val = val - 10
			carry = true
		} else {
			carry = false
		}
		n := &ListNode{
			Val:  val,
			Next: nil,
		}

		next.Next = n
		next = n

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		next.Next = l1
	}
	if l2 != nil {
		next.Next = l2
	}
	for next.Next != nil && carry {
		next = next.Next
		if carry {
			val = next.Val + 1
		} else {
			val = next.Val
		}

		if val >= 10 {
			val = val - 10
			carry = true
		} else {
			carry = false
		}
		next.Val = val
	}

	if carry {
		next.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}

	return fake.Next
}
