package RotateList

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	countNode := head
	firstNode := head
	num := 1
	for countNode.Next != nil {
		num++
		countNode = countNode.Next
	}
	countNode.Next = firstNode
	countNode = countNode.Next

	k = k % num

	for k > 0 {
		countNode = countNode.Next
		k--
	}

	firstNode = countNode.Next
	countNode.Next = nil
	return firstNode
}
