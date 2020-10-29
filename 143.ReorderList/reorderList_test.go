package ReorderList

import (
	"testing"
)

func Test_reorderList(t *testing.T) {
	tests := []struct {
		head *ListNode
	}{
		{
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							nil,
						},
					},
				},
			},
		},
		{
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
		},
	}

	for _, testItem := range tests {
		reorderList(testItem.head)
	}

}
