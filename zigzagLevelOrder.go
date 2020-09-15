package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		val := make([]int, 0)
		var kids []*TreeNode

		for _, node := range queue {
			if level%2 == 0 {
				val = append(val, node.Val)
			} else {
				val = append([]int{node.Val}, val...)

			}

			if node.Left != nil {
				kids = append(kids, node.Left)
			}
			if node.Right != nil {
				kids = append(kids, node.Right)
			}
		}
		level++
		result = append(result, val)
		queue = kids
	}
	return result

}
