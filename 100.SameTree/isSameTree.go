package SameTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if p.Val != q.Val {
		return false
	}

	isSame := isSameTree(p.Left, q.Left)
	if !isSame {
		return false
	}
	isSame = isSameTree(p.Right, q.Right)
	return isSame
}

// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	l := isSameTreeHelper(p)
// 	r := isSameTreeHelper(q)
// 	if len(l) != len(r) {
// 		return false
// 	}

// 	for index, val := range l {
// 		if val != r[index] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func isSameTreeHelper(root *TreeNode) []int {
// 	var result []int
// 	if root == nil {
// 		result = append(result, 0)
// 		return result
// 	}
// 	result = append(result, root.Val)
// 	result = append(result, isSameTreeHelper(root.Left)...)
// 	result = append(result, isSameTreeHelper(root.Right)...)

// 	return result
// }
