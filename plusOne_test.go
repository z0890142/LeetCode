package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_plusOne(t *testing.T) {

	tests := []struct {
		digits []int
		result []int
		Name   string
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 4},
			"test 1",
		},
		{
			[]int{9, 9, 9},
			[]int{1, 0, 0, 0},
			"test 2",
		},
	}

	for _, testItem := range tests {
		Convey(testItem.Name, t, func() {
			So(plusOne(testItem.digits), ShouldResemble, testItem.result)
		})

	}
}
