package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextNode := head.Next

	if nextNode != nil {
		head.Next = nextNode.Next
		nextNode.Next = head
		head.Next = swapPairs(head.Next)
	}

	return nextNode
}
