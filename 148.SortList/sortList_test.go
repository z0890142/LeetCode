package SortList

import (
	"fmt"
	"testing"
)

func Test_sortList(t *testing.T) {
	tests := []struct {
		head *ListNode
	}{
		{
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						2,
						&ListNode{
							1,
							nil,
						},
					},
				},
			},
		},
	}

	for _, testItem := range tests {
		t := sortList(testItem.head)
		fmt.Println(t)
	}

}
