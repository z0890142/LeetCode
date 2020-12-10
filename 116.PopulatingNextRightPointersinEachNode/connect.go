package PopulatingNextRightPointersinEachNode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	level := root
	cur := level
	for level != nil {
		cur = level
		for cur != nil {
			if cur.Left != nil {
				cur.Left.Next = cur.Right
			}
			if cur.Right != nil && cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}

			cur = cur.Next
		}
		level = level.Left

	}
	return root
}
