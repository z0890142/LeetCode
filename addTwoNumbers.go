package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultNode ListNode
	if l1.Next == nil && l2.Next == nil {
		return nil
	}
	resultNode.Val = (l1.Val + l2.Val) % 10

	tmpVal := (l1.Val + l2.Val) / 10
	if l1.Next == nil {
		l1.Next = &ListNode{Val: tmpVal, Next: nil}
	} else {
		l1.Next.Val += tmpVal
	}
	resultNode.Next = addTwoNumbers(l1.Next, l2.Next)
	return &resultNode
}
