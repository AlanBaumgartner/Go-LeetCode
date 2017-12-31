package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	n := &ListNode{l1.Val + l2.Val, addTwoNumbers(l1.Next, l2.Next)}
	if n.Val > 9 {
		n = &ListNode{n.Val - 10, addTwoNumbers(n.Next, &ListNode{1, nil})}
	}
	return n
}
