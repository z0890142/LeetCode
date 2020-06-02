package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := ListNode{}

	if l2.Val <= l1.Val {
		result.Val = l2.Val
		result.Next = mergeTwoLists(l1, l2.Next)
	} else if l1.Val <= l2.Val {
		result.Val = l1.Val
		result.Next = mergeTwoLists(l2, l1.Next)

	}
	return &result
}
