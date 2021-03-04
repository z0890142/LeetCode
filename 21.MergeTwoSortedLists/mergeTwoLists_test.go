package MergeTwoSortedLists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mergeTwoListstest1 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	},
}
var mergeTwoListstest2 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	},
}

var mergeTwoListstest3 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	},
}
var mergeTwoListstest4 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	},
}
var mergeTwoListstest5 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	},
}

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		l1     *ListNode
		l2     *ListNode
		result []int
	}{
		{
			&mergeTwoListstest1,
			&mergeTwoListstest2,
			[]int{1, 1, 2, 3, 4, 4},
		}, {
			&mergeTwoListstest3,
			nil,
			[]int{1, 2, 4},
		}, {
			&mergeTwoListstest4,
			&mergeTwoListstest5,
			[]int{1, 1, 3, 4, 5, 6},
		},
	}
	for _, testItem := range tests {
		result := mergeTwoLists(testItem.l1, testItem.l2)
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
