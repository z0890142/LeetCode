package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var deleteDuplicates2Liststest1 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
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
	},
}
var deleteDuplicates2Liststest2 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	},
}

var deleteDuplicates2Liststest3 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	},
}

func Test_deleteDuplicates2(t *testing.T) {
	tests := []struct {
		head   *ListNode
		result []int
	}{
		{
			&deleteDuplicates2Liststest1,
			[]int{1, 2, 4, 5},
		}, {
			&deleteDuplicates2Liststest2,
			[]int{1, 2, 3},
		}, {
			&deleteDuplicates2Liststest3,
			[]int{2, 3},
		},
	}
	for _, testItem := range tests {
		result := deleteDuplicates2(testItem.head)
		resultList := []int{}
		getNode(result, &resultList)
		assert.Equal(t, testItem.result, resultList)

	}

}
