package InsertionSortList

import "testing"

func Test_insertionSortList(t *testing.T) {
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
		insertionSortList(testItem.head)
	}

}
