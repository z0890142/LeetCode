package BinaryTreePreorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	left := preorderTraversal(root.Left)
	rigth := preorderTraversal(root.Right)
	result = append(result, root.Val)

	result = append(result, left...)
	result = append(result, rigth...)
	return result
}
