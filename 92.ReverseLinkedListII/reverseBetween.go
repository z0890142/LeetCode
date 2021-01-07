package ReverseLinkedListII

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	dummy := ListNode{
		0,
		head,
	}
	prev := &dummy
	cur := &dummy
	var then *ListNode

	for i := 0; i < m; i++ {
		prev = cur
		cur = cur.Next
		then = cur.Next
	}
	for i := 0; i < n-m; i++ {
		cur.Next = then.Next
		then.Next = prev.Next
		prev.Next = then
		then = cur.Next
	}

	return dummy.Next
}
