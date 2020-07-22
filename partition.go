package main

func partition(head *ListNode, x int) *ListNode {
	smallFirstNode := &ListNode{}
	bigFirstNode := &ListNode{}
	smallNode := smallFirstNode
	bigNode := bigFirstNode

	for head != nil {
		if head.Val >= x {
			bigNode.Next = head
			bigNode = bigNode.Next
		} else {
			smallNode.Next = head
			smallNode = smallNode.Next
		}
		head = head.Next
	}
	bigNode.Next = nil
	smallNode.Next = bigFirstNode.Next

	return smallFirstNode.Next
}
