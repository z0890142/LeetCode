package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left, right int
	if root.Left != nil {
		left = maxDepth(root.Left)
	}
	if root.Right != nil {
		right = maxDepth(root.Left)
	}
	if right > left {
		return right + 1
	} else {
		return left + 1
	}
}

func maxDepthHelper(root *TreeNode, level int) int {
	if root == nil {
		return level - 1
	}
	var left, right int
	if root.Left != nil {
		left = maxDepthHelper(root.Left, level+1)
	}
	if root.Right != nil {
		right = maxDepthHelper(root.Left, level+1)
	}

	if right > left {
		return right
	} else {
		return left
	}
}
