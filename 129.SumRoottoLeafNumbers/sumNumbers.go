package SumRoottoLeafNumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var result int
	sumNumbersHelper(root, 0, &result)
	return result
}

func sumNumbersHelper(root *TreeNode, tmp int, result *int) {
	if root == nil {
		return
	}
	tmp = 10*tmp + root.Val

	if root.Right == nil && root.Left == nil {
		*result += tmp
		return
	}
	sumNumbersHelper(root.Left, tmp, result)
	sumNumbersHelper(root.Right, tmp, result)

}

// func sumNumbers(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	r, l := 0, 0
// 	if root.Left != nil {
// 		root.Left.Val = root.Left.Val + root.Val*10
// 		l = sumNumbers(root.Left)
// 	}
// 	if root.Right != nil {
// 		root.Right.Val = root.Right.Val + root.Val*10
// 		r = sumNumbers(root.Right)
// 	}
// 	if r == 0 && l == 0 {
// 		return root.Val
// 	}

// 	return r + l
// }
