package InsertionSortList

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	pseudo_head := &ListNode{Next: nil}
	curr := head
	for curr != nil {

		prev_node := pseudo_head
		next_node := prev_node.Next

		for next_node != nil {
			if curr.Val < next_node.Val {
				break
			}
			prev_node = next_node
			next_node = next_node.Next
		}
		next_iter := curr.Next

		prev_node.Next = curr
		curr.Next = next_node

		curr = next_iter

	}
	return pseudo_head.Next
}
