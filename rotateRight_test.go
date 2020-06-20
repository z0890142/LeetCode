package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rtest1 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	},
}
var rtest2 = ListNode{
	Val: 0,
	Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	},
}

func Test_rotateRight(t *testing.T) {
	tests := []struct {
		head   *ListNode
		k      int
		result []int
	}{
		{
			&rtest1,
			2,
			[]int{4, 5, 1, 2, 3},
		},
		{
			&rtest2,
			4,
			[]int{2, 0, 1},
		},
	}

	for _, testItem := range tests {
		result := rotateRight(testItem.head, testItem.k)
		resultList := []int{}
		getNode(result, &resultList)
		assert.Equal(t, testItem.result, resultList)

	}
}
