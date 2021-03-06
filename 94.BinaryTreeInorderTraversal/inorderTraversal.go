package BinaryTreeInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result

}

// func inorderTraversal(root *TreeNode) []int {
// 	var res []int
// 	if root == nil {
// 		return res
// 	}
// 	stack := []*TreeNode{}
// 	for root != nil || len(stack) > 0 {
// 		for ; root != nil; root = root.Left {
// 			stack = append(stack, root)
// 		}
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		res = append(res, node.Val)
// 		if node.Right != nil {
// 			root = node.Right
// 		}
// 	}
// 	return res
// }
