package ConvertSortedListtoBinarySearchTree

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var tmp *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var size int
	tmp = head
	for head != nil {
		size++
		head = head.Next
	}

	return convertToBST(0, size-1)
}

func convertToBST(start int, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	left := convertToBST(start, mid-1)
	node := &TreeNode{Val: tmp.Val}
	node.Left = left

	tmp = tmp.Next
	node.Right = convertToBST(mid+1, end)
	return node
}

// func sortedListToBST(head *ListNode) *TreeNode {
// 	var result TreeNode
// 	if head == nil {
// 		return nil
// 	}
// 	mid := findMidList(head)
// 	result.Val = mid.Val
// 	if head == mid {
// 		return &result
// 	}
// 	result.Right = sortedListToBST(mid.Next)
// 	result.Left = sortedListToBST(head)
// 	return &result

// }

// func findMidList(head *ListNode) *ListNode {
// 	prevPtr := head
// 	slowPtr := head
// 	fastPtr := head
// 	for fastPtr != nil && fastPtr.Next != nil {
// 		prevPtr = slowPtr
// 		fastPtr = fastPtr.Next.Next
// 		slowPtr = slowPtr.Next
// 	}

// 	prevPtr.Next = nil

// 	return slowPtr
// }
