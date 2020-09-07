package main

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
