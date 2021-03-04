package MergeTwoSortedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	n := root
	for !(l1 == nil && l2 == nil) {
		if l1 != nil && (l2 == nil || l2.Val > l1.Val) {
			n.Next = l1
			l1 = l1.Next
		} else {
			n.Next = l2
			l2 = l2.Next
		}
		n = n.Next
	}
	return root.Next
}

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}
// 	if l2 == nil {
// 		return l1
// 	}
// 	result := ListNode{}

// 	if l2.Val <= l1.Val {
// 		result.Val = l2.Val
// 		result.Next = mergeTwoLists(l1, l2.Next)
// 	} else if l1.Val <= l2.Val {
// 		result.Val = l1.Val
// 		result.Next = mergeTwoLists(l2, l1.Next)

// 	}
// 	return &result
// }
