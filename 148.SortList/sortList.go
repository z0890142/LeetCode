package SortList

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMidList(head)

	left := sortList(head)
	right := sortList(mid)
	return merge(left, right)
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummyHead := &ListNode{}
	head := dummyHead
	for left != nil && right != nil {
		if left.Val < right.Val {
			head.Next = left
			left = left.Next
		} else {
			head.Next = right
			right = right.Next
		}
		head = head.Next

	}
	if left != nil {
		head.Next = left
	} else {
		head.Next = right

	}
	return dummyHead.Next
}

func findMidList(head *ListNode) *ListNode {
	prevPtr := head
	slowPtr := head
	fastPtr := head
	for fastPtr != nil && fastPtr.Next != nil {
		prevPtr = slowPtr
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
	}

	prevPtr.Next = nil

	return slowPtr
}
