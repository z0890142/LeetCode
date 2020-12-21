package ConstructBinaryTreefromPreorderandInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var root int
	node := TreeNode{Val: preorder[0]}
	for index, val := range inorder {
		if preorder[0] == val {
			root = index
			break
		}
	}
	node.Left = buildTree(preorder[1:root], inorder[0:root])
	node.Right = buildTree(preorder[root+1:], inorder[root+1:])

	return &node
}
