package ReverseLinkedListII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var reverseBetweenListstest1 = ListNode{
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

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		head   *ListNode
		m      int
		n      int
		result []int
	}{
		{
			&reverseBetweenListstest1,
			2,
			4,
			[]int{1, 4, 3, 2, 5},
		},
	}
	for _, testItem := range tests {
		result := reverseBetween(testItem.head, testItem.m, testItem.n)
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
