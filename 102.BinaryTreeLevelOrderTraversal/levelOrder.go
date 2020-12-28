package BinaryTreeLevelOrderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
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

			val = append(val, node.Val)

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

// func levelOrder(root *TreeNode) [][]int {
// 	var result [][]int
// 	if root == nil {
// 		return result
// 	}
// 	var tmpResult map[int][]int
// 	tmpResult = make(map[int][]int)
// 	tmpResult[0] = []int{root.Val}

// 	levelOrderHelper(root.Left, 1, &tmpResult)
// 	levelOrderHelper(root.Right, 1, &tmpResult)

// 	for i := 0; i < len(tmpResult); i++ {
// 		result = append(result, tmpResult[i])
// 	}

// 	return result
// }

// func levelOrderHelper(root *TreeNode, level int, tmpResult *map[int][]int) {
// 	if root == nil {
// 		return
// 	}
// 	(*tmpResult)[level] = append((*tmpResult)[level], root.Val)

// 	levelOrderHelper(root.Left, level+1, tmpResult)
// 	levelOrderHelper(root.Right, level+1, tmpResult)
// 	return
// }
