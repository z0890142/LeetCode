package main

func isBalanced(root *TreeNode) bool {
	_, res := isBalancedHelper(root)
	return res
}

func isBalancedHelper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	l, isBall := isBalancedHelper(root.Left)
	r, isBalr := isBalancedHelper(root.Right)
	if !(isBall && isBalr) {
		return 0, false
	}
	if Abs(l-r) > 1 {
		return 0, false
	} else {
		return maxdp(l, r), true
	}
}

func maxdp(l int, r int) int {
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
