package KthSmallestElementisaBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func kthSmallest(root *TreeNode, k int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	count := kthSmallestHelper(root.Left)
// 	if k < count+1 {
// 		return kthSmallest(root.Left, k)
// 	}
// 	if k > count+1 {
// 		return kthSmallest(root.Right, k-count-1)
// 	}
// 	return root.Val
// }

// func kthSmallestHelper(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return 1 + kthSmallestHelper(root.Left) + kthSmallestHelper(root.Right)
// }

func kthSmallest(root *TreeNode, k int) int {
	cur := 0
	return kthSmallestHelper(root, k, &cur)
}

func kthSmallestHelper(root *TreeNode, k int, cur *int) int {
	if root == nil {
		return -1
	}

	leftRet := kthSmallestHelper(root.Left, k, cur)
	if leftRet != -1 {
		return leftRet
	}
	if *cur == k-1 {
		return root.Val
	}
	*cur = *cur + 1
	rightRet := kthSmallestHelper(root.Right, k, cur)
	return rightRet
}
