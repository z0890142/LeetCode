package LinkedListCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head.Next
	fast := head.Next.Next
	if fast == nil {
		return false
	}

	for slow != fast {
		if fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next

		if slow == nil || fast == nil {
			return false
		}
	}
	return true
}
