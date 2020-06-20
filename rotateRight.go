package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	countNode := head
	firstNode := head

	nums := 0
	for true {
		nums += 1
		if countNode.Next == nil {
			countNode.Next = firstNode
			countNode = countNode.Next
			break
		}
		countNode = countNode.Next
	}
	count := nums - (k % nums)

	for count >= 1 {
		if count == 1 {
			firstNode = countNode.Next
			countNode.Next = nil
			break
		}
		countNode = countNode.Next
		count -= 1
	}
	return firstNode
}
