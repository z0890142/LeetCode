package main

func hasPathSum(root *TreeNode, sum int) bool {
	if sum == 0 {
		return true
	}
	if root == nil {
		return false
	}

	sum -= root.Val
	left, right := hasPathSum(root.Left, sum), hasPathSum(root.Right, sum)
	return left || right
}
