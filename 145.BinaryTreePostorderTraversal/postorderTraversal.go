package BinaryTreePostorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	left := postorderTraversal(root.Left)
	rigth := postorderTraversal(root.Right)

	result = append(result, left...)
	result = append(result, rigth...)
	result = append(result, root.Val)

	return result
}
