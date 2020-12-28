package SymmetricTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	left := isSymmetricHelperLeft(root.Left)
	right := isSymmetricHelper(root.Right)

	if len(left) != len(right) {
		return false
	}
	for index, val := range left {
		if val != right[index] {
			return false
		}
	}
	return true
}

func isSymmetricHelperLeft(root *TreeNode) []int {
	var result []int

	if root == nil {
		result = append(result, 0)
		return result
	}
	result = append(result, root.Val)
	result = append(result, isSymmetricHelperLeft(root.Right)...)
	result = append(result, isSymmetricHelperLeft(root.Left)...)
	return result
}

func isSymmetricHelper(root *TreeNode) []int {
	var result []int

	if root == nil {
		result = append(result, 0)
		return result
	}
	result = append(result, root.Val)
	result = append(result, isSymmetricHelper(root.Left)...)
	result = append(result, isSymmetricHelper(root.Right)...)
	return result
}
