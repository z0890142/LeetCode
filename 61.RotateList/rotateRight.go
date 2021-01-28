package RotateList

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
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

	num = num - (k % num)

	for i := 1; i < num; i++ {
		countNode = countNode.Next
	}
	firstNode = countNode.Next
	countNode.Next = nil

	return firstNode
}
