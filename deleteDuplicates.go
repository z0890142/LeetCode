package main

func deleteDuplicates(head *ListNode) *ListNode {
	firstNode := head

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}
	return firstNode
}
