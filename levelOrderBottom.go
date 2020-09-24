package main

func levelOrderBottom(root *TreeNode) [][]int {
	rs := make([][]int, 0)

	if root == nil {
		return rs
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var line []int
		var kids []*TreeNode

		for _, node := range queue {
			line = append(line, node.Val)

			if node.Left != nil {
				kids = append(kids, node.Left)
			}
			if node.Right != nil {
				kids = append(kids, node.Right)
			}
		}

		rs = append([][]int{line}, rs...)
		queue = kids
	}

	return rs
}
