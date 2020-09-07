package main

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var isBST bool
	isBST = isValidBSTHelper(root, math.MinInt64, math.MaxInt64)

	return isBST
}
func isValidBSTHelper(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}

	return isValidBSTHelper(root.Left, min, root.Val) && isValidBSTHelper(root.Right, root.Val, max)
}
