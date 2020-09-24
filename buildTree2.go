package main

//Construct Binary Tree from Preorder and Inorder Traversal
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var root int
	node := TreeNode{Val: postorder[len(postorder)-1]}
	for index, val := range inorder {
		if postorder[len(postorder)-1] == val {
			root = index
			break
		}
	}
	node.Left = buildTree(inorder[0:root], postorder[0:root])
	node.Right = buildTree(inorder[root+1:], postorder[root:len(postorder)-1])

	return &node
}
