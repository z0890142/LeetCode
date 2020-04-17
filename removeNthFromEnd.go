package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = removeNode(head, &n)
	return head
}

func removeNode(node *ListNode, n *int) *ListNode {
	if node.Next == nil {
		*n -= 1
		if *n == 0 {
			return node.Next
		}
		return node
	}
	node.Next = removeNode(node.Next, n)
	*n -= 1
	if *n == 0 {
		return node.Next
	}
	return node
}
