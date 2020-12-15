package FlattenBinaryTreetoLinkedList

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	mainParent := root
	for mainParent != nil {
		if mainParent.Right == nil {
			mainParent.Right = mainParent.Left
			mainParent.Left = nil
		} else {
			tempFindLeaf := mainParent.Left
			if tempFindLeaf != nil {
				for tempFindLeaf.Right != nil {
					tempFindLeaf = tempFindLeaf.Right
				}
				tempFindLeaf.Right = mainParent.Right
				mainParent.Right = mainParent.Left
				mainParent.Left = nil
			}
		}
		mainParent = mainParent.Right
	}
}

// DFS
// func flatten(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	flattenHelper(root)
// 	fmt.Println(root)
// }

// func flattenHelper(root *TreeNode) *TreeNode {

// 	if root == nil {
// 		return nil
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return root
// 	}

// 	left := flattenHelper(root.Left)
// 	right := flattenHelper(root.Right)
// 	if left != nil {
// 		root.Right = left
// 		for left.Right != nil {
// 			left = left.Right
// 		}
// 		left.Right = right
// 		root.Left = nil
// 	} else {
// 		root.Right = right
// 	}
// 	return root
// }
