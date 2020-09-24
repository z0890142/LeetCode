package main

//Construct Binary Tree from Preorder and Inorder Traversal
func buildTree1(preorder []int, inorder []int) *TreeNode {
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
	node.Left = buildTree1(preorder[1:root+1], inorder[0:root])
	node.Right = buildTree1(preorder[root+1:], inorder[root+1:])

	return &node
}
