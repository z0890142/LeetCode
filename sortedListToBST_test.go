package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sortedListToBSTtest1 = ListNode{
	Val: -10,
	Next: &ListNode{
		Val: -3,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	},
}

func Test_sortedListToBST(t *testing.T) {
	tests := []struct {
		head   *ListNode
		result *TreeNode
	}{
		{
			&sortedListToBSTtest1,
			&TreeNode{
				0,
				&TreeNode{
					-3,
					&TreeNode{
						-10,
						nil,
						nil,
					},
					nil,
				},
				&TreeNode{
					9,
					&TreeNode{
						5,
						nil,
						nil,
					},
					nil,
				},
			},
		},
	}
	for _, testItem := range tests {
		result := sortedListToBST(testItem.head)
		assert.Equal(t, testItem.result, result)

	}

}
