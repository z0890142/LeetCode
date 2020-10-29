package CopyRandomList

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	result := make(map[*Node]*Node)

	return Helper(head, result)
}

func Helper(head *Node, result map[*Node]*Node) *Node {
	if head == nil {
		return nil
	}

	if val, isExist := result[head]; isExist {
		return val
	}
	newHead := &Node{Val: head.Val}
	result[head] = newHead
	newHead.Next = Helper(head.Next, result)
	newHead.Random = Helper(head.Random, result)
	return newHead
}
