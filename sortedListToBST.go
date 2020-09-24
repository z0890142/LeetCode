package main

func sortedListToBST(head *ListNode) *TreeNode {
	var result TreeNode
	if head == nil {
		return nil
	}
	mid := findMidList(head)
	result.Val = mid.Val
	if head == mid {
		return &result
	}
	result.Right = sortedListToBST(mid.Next)
	result.Left = sortedListToBST(head)
	return &result
}

func findMidList(head *ListNode) *ListNode {
	prevPtr := head
	slowPtr := head
	fastPtr := head
	for fastPtr != nil && fastPtr.Next != nil {
		prevPtr = slowPtr
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
	}

	prevPtr.Next = nil

	return slowPtr
}
