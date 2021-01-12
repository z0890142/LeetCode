package PartitionList

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	bigFirstHead := &ListNode{}
	smallFirstHead := &ListNode{}
	smallNode := smallFirstHead
	bigNode := bigFirstHead

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
	smallNode.Next = bigFirstHead.Next
	return smallFirstHead.Next
}
