package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test1 = ListNode{
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
var test3 = ListNode{
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

var test2 = ListNode{
	Val:  1,
	Next: nil,
}

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		head   *ListNode
		n      int
		result []int
	}{{
		&test1,
		2,
		[]int{1, 2, 3, 5},
	}, {
		&test3,
		1,
		[]int{1, 2, 3, 4},
	}, {
		&test2,
		1,
		[]int{},
	},
	}
	for _, testItem := range tests {
		result := removeNthFromEnd(testItem.head, testItem.n)
		resultList := []int{}
		getNode(result, &resultList)
		assert.Equal(t, testItem.result, resultList)

	}

}

func getNode(node *ListNode, preList *[]int) {
	if node == nil {
		return
	}
	*preList = append(*preList, node.Val)
	if node.Next == nil {
		return
	}
	getNode(node.Next, preList)
}
