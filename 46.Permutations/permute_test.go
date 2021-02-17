package Permutations

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_permute(t *testing.T) {

	tests := []struct {
		nums   []int
		result [][]int
		Name   string
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
				[]int{2, 1, 3},
				[]int{2, 3, 1},
				[]int{3, 1, 2},
				[]int{3, 2, 1},
			},
			"test 1",
		},
	}

	for _, testItem := range tests {
		Convey(testItem.Name, t, func() {
			So(permute(testItem.nums), ShouldEqual, testItem.result)
		})

	}
}
