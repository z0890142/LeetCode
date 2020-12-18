package MinimumDepthofBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}

	if left > right {
		return right + 1
	} else {
		return left + 1
	}
}

// func minDepth(root *TreeNode) int {
// 	return minDepthHelper(root, 0)
// }

// func minDepthHelper(root *TreeNode, count int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return count + 1
// 	}
// 	var l, r int
// 	l = minDepthHelper(root.Left, count+1)
// 	r = minDepthHelper(root.Right, count+1)

// 	return min(l, r)
// }

// func min(a int, b int) int {
// 	if a == 0 || b == 0 {
// 		return a + b
// 	}
// 	if a > b {
// 		return b
// 	}
// 	return a
// }
