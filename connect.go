package main

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
	nodeList := []*Node{root}
	for len(nodeList) > 0 {
		var tmpList []*Node
		if nodeList[0].Left != nil {
			tmpList = append(tmpList, nodeList[0].Left)
		}
		if nodeList[0].Right != nil {
			tmpList = append(tmpList, nodeList[0].Right)
		}
		for index := 1; index < len(nodeList); index++ {
			nodeList[index-1].Next = nodeList[index]
			if nodeList[index].Left != nil {
				tmpList = append(tmpList, nodeList[index].Left)
			}
			if nodeList[index].Right != nil {
				tmpList = append(tmpList, nodeList[index].Right)
			}
		}
		nodeList = tmpList
	}
	return root
}
