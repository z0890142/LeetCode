package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_uniquePaths(t *testing.T) {

	tests := []struct {
		m      int
		n      int
		result int
		Name   string
	}{
		{
			3,
			2,
			3,
			"test 1",
		},
		{
			7,
			3,
			28,
			"test 1",
		},
	}

	for _, testItem := range tests {
		Convey(testItem.Name, t, func() {
			So(uniquePaths(testItem.m, testItem.n), ShouldEqual, testItem.result)
		})

	}
}
