package main

func deleteDuplicates2(head *ListNode) *ListNode {
	firstNode := ListNode{Val: -9999, Next: head}
	firstNode.Next = deleteDuplicates2Helper(firstNode.Next, firstNode.Val)
	return firstNode.Next
}

func deleteDuplicates2Helper(head *ListNode, preVal int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == preVal {
		return deleteDuplicates2Helper(head.Next, head.Val)
	}
	if head.Next != nil && head.Val == head.Next.Val {
		return deleteDuplicates2Helper(head.Next.Next, head.Val)
	}
	head.Next = deleteDuplicates2Helper(head.Next, head.Val)
	return head
}
