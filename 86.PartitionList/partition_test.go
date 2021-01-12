package PartitionList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var partitionList1 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
		},
	},
}
var partitionList2 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	},
}

var partitionList3 = ListNode{
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

func Test_partition(t *testing.T) {
	tests := []struct {
		head   *ListNode
		x      int
		result []int
	}{
		{
			&partitionList1,
			3,
			[]int{1, 2, 2, 4, 3, 5},
		},
	}
	for _, testItem := range tests {
		result := partition(testItem.head, testItem.x)
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
