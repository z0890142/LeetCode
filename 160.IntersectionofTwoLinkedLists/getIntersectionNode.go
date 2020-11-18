package IntersectionofTwoLinkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

// 共同的node: Common
// listA = A + Common
// listB =B + Common
// PointerA: [A + Common + B] + Common
// PointerB: [B + Common + A] + Common

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	if p1 == nil || p2 == nil {
		return nil
	}
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
		if p1 == p2 {
			return p1
		}
		if p1 == nil {
			p1 = headB
		}
		if p2 == nil {
			p2 = headA
		}
	}
	return p1
}
