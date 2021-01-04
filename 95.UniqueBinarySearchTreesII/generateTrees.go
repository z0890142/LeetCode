package UniqueBinarySearchTreesII

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	if n == 1 {
		return []*TreeNode{
			{
				Val: 1,
			},
		}
	}
	var result []*TreeNode
	result = generateTreesHelper(1, n)
	return result
}

func generateTreesHelper(min int, max int) (result []*TreeNode) {
	if min > max {
		return []*TreeNode{nil}
	}
	for i := min; i <= max; i++ {
		leftSubTree := generateTreesHelper(min, i-1)
		rightSubTree := generateTreesHelper(i+1, max)
		fmt.Println(len(leftSubTree))
		for _, leftTree := range leftSubTree {
			for _, rightTree := range rightSubTree {
				root := TreeNode{Val: i}
				root.Left = leftTree
				root.Right = rightTree
				result = append(result, &root)
			}
		}
	}

	return
}
