package PermutationsII

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_permuteUnique(t *testing.T) {

	tests := []struct {
		nums   []int
		result [][]int
		Name   string
	}{
		{
			[]int{1, 1, 2},
			[][]int{
				[]int{1, 1, 2},
				[]int{1, 2, 1},
				[]int{2, 1, 1},
			},
			"test 1",
		},
	}

	for _, testItem := range tests {
		Convey(testItem.Name, t, func() {
			So(permuteUnique(testItem.nums), ShouldEqual, testItem.result)
		})

	}
}
