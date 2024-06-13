package p2

func main() {
	l1 := []int{9, 9, 9, 9}
	l2 := []int{9, 9, 9, 9, 9, 9, 9}
	addTwoNumbersReverse(makeList(l1), makeList(l2))
}

func makeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{}
	fake := &ListNode{}
	root = fake
	for _, num := range nums {
		n := &ListNode{
			Val:  num,
			Next: nil,
		}
		root.Next = n
		root = n
	}
	return fake.Next
}

// 可视为原题的扩展
func addTwoNumbersReverse(l1 *ListNode, l2 *ListNode) *ListNode {
	var list1, list2 []*ListNode
	for l1 != nil {
		list1 = append(list1, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		list2 = append(list2, l2)
		l2 = l2.Next
	}

	if len(list1) > len(list2) {
		return addTwoReverse(list2, list1)
	} else {
		return addTwoReverse(list1, list2)
	}
}

func addTwoReverse(list1 []*ListNode, list2 []*ListNode) *ListNode {
	carry := false
	val := 0
	for i := len(list1) - 1; i >= 0; i-- {
		if carry {
			val = list1[i].Val + list2[i].Val + 1
		} else {
			val = list1[i].Val + list2[i].Val
		}

		if val >= 10 {
			val = val - 10
			carry = true
		} else {
			carry = false
		}
		list2[i].Val = val
	}

	if len(list2) > len(list1) {
		jointNodeIdx := len(list2) - len(list1) - 1
		for j := jointNodeIdx; carry && j >= 0; j-- {
			if carry {
				val = list2[j].Val + 1
			} else {
				val = list2[j].Val
			}
			if val >= 10 {
				val = val - 10
				carry = true
				list2[j].Val = val
			} else {
				carry = false
			}
		}
	}

	if carry {
		root := &ListNode{
			Val:  1,
			Next: list2[0],
		}
		return root
	} else {
		return list2[0]
	}
}
