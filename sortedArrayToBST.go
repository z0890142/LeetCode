package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := new(TreeNode)
	middle := len(nums) / 2
	root.Val = nums[middle]
	if middle != 0 {
		root.Left = sortedArrayToBST(nums[:middle])
		root.Right = sortedArrayToBST(nums[middle+1:])
	}
	return root
}
