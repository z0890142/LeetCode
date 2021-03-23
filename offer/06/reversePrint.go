package reversePrint

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var prev *ListNode = nil
	l := 0
	for head != nil {
		l++
		tmpNext := head.Next
		head.Next = prev
		prev = head
		head = tmpNext
	}

	result := make([]int, l)
	l = 0
	for prev != nil {
		result[l] = prev.Val
		prev = prev.Next
	}
	return result
}
