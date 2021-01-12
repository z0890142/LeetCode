package RemoveDuplicatesfromSortedListII

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	firstNode := ListNode{Val: -9999, Next: head}
	firstNode.Next = deleteDuplicatesHelper(firstNode.Next)
	return firstNode.Next
}

func deleteDuplicatesHelper(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicatesHelper(head.Next)
	}

	head.Next = deleteDuplicatesHelper(head.Next)
	return head
}
