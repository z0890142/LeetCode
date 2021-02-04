package BinaryTreeRightSideView

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//DFS
func rightSideView(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	rightSideViewHelper(root, 0, &result)

	return result
}

func rightSideViewHelper(root *TreeNode, level int, result *[]int) {
	if root == nil {
		return
	}

	if len(*result) == level {
		*result = append(*result, root.Val)
	}
	rightSideViewHelper(root.Right, level+1, result)
	rightSideViewHelper(root.Left, level+1, result)

}

// BFS
// func rightSideView(root *TreeNode) []int {
// 	var result []int
// 	queue := []*TreeNode{root}

// 	if root == nil {
// 		return result
// 	}

// 	for len(queue) > 0 {
// 		var tmp []*TreeNode

// 		for index, node := range queue {
// 			if node.Left != nil {
// 				tmp = append(tmp, node.Left)
// 			}
// 			if node.Right != nil {
// 				tmp = append(tmp, node.Right)
// 			}
// 			if index == len(queue)-1 {
// 				result = append(result, node.Val)
// 			}
// 		}
// 		queue = tmp
// 	}
// 	return result
// }
