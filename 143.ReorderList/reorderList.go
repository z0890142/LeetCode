package ReorderList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	var fast, slow *ListNode
	fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	curr := slow
	var prev, next *ListNode
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	first := head
	second := prev

	for second.Next != nil {
		next = first.Next
		first.Next = second
		first = next

		next = second.Next
		second.Next = first
		second = next
	}

}
