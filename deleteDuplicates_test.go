package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var deleteDuplicatesListstest1 = ListNode{
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
var deleteDuplicatesListstest2 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	},
}

var deleteDuplicatesListstest3 = ListNode{
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

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		head   *ListNode
		result []int
	}{
		{
			&deleteDuplicatesListstest1,
			[]int{1, 2, 4, 5},
		}, {
			&deleteDuplicatesListstest2,
			[]int{1, 2, 3},
		}, {
			&deleteDuplicatesListstest3,
			[]int{2, 3},
		},
	}
	for _, testItem := range tests {
		result := deleteDuplicates(testItem.head)
		resultList := []int{}
		getNode(result, &resultList)
		assert.Equal(t, testItem.result, resultList)

	}

}
