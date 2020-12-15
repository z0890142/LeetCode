package PathSumII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	pathSumHelper(root, sum, []int{}, &result)
	return result
}

func pathSumHelper(root *TreeNode, sum int, tmpPath []int, pathSumResult *[][]int) {
	sum = sum - root.Val
	tmpPath = append(tmpPath, root.Val)
	if root.Left == nil && root.Right == nil {

		if sum == 0 {
			tmp := make([]int, len(tmpPath))
			tmp = append([]int{}, tmpPath...)
			*pathSumResult = append(*pathSumResult, tmp)
			return
		}
		return
	}

	if root.Left != nil {
		pathSumHelper(root.Left, sum, tmpPath, pathSumResult)
	}
	if root.Right != nil {
		pathSumHelper(root.Right, sum, tmpPath, pathSumResult)
	}
	return

}

// func pathSum(root *TreeNode, sum int) [][]int {
// 	var pathSumResult [][]int

// 	if root == nil {
// 		return pathSumResult
// 	}

// 	if root.Left == nil && root.Right == nil {
// 		if sum == root.Val {
// 			return append(pathSumResult, []int{root.Val})
// 		}
// 		return pathSumResult
// 	}

// 	for _, path := range pathSum(root.Left, sum-root.Val) {
// 		pathSumResult = append(pathSumResult, append([]int{root.Val}, path...))
// 	}
// 	for _, path := range pathSum(root.Right, sum-root.Val) {
// 		pathSumResult = append(pathSumResult, append([]int{root.Val}, path...))
// 	}
// 	return pathSumResult
// }
