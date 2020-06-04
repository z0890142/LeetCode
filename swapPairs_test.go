package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var swapPairstest1 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	},
}

var swapPairstest2 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	},
}

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		head   *ListNode
		result []int
	}{
		{
			&swapPairstest1,
			[]int{2, 1, 4, 3},
		}, {
			&swapPairstest2,
			[]int{2, 1, 3, 4},
		},
	}
	for _, testItem := range tests {
		result := swapPairs(testItem.head)
		resultList := []int{}
		getNode(result, &resultList)
		assert.Equal(t, testItem.result, resultList)

	}

}
