package main

func deleteDuplicates(head *ListNode) *ListNode {
	firstNode := ListNode{Val: -9999, Next: head}
	firstNode.Next = deleteDuplicatesHelper(firstNode.Next, firstNode.Val)
	return firstNode.Next
}

func deleteDuplicatesHelper(head *ListNode, preVal int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == preVal {
		return deleteDuplicatesHelper(head.Next, head.Val)
	}
	if head.Next != nil && head.Val == head.Next.Val {
		return deleteDuplicatesHelper(head.Next.Next, head.Val)
	}
	head.Next = deleteDuplicatesHelper(head.Next, head.Val)
	return head
}
