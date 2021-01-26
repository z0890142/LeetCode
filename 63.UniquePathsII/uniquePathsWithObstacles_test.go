package UniquePathsII

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_uniquePathsWithObstacles(t *testing.T) {

	tests := []struct {
		obstacleGrid [][]int
		result       int
		Name         string
	}{
		// {
		// 	[][]int{
		// 		[]int{0, 0, 0},
		// 		[]int{0, 1, 0},
		// 		[]int{0, 0, 0},
		// 	},
		// 	2,
		// 	"test 1",
		// },
		// {
		// 	[][]int{
		// 		[]int{1},
		// 	},
		// 	0,
		// 	"test 2",
		// },
		{
			[][]int{
				[]int{0},
			},
			1,
			"test 3",
		},
		{
			[][]int{
				[]int{1, 0},
			},
			0,
			"test 4",
		},
	}

	for _, testItem := range tests {
		Convey(testItem.Name, t, func() {
			So(uniquePathsWithObstacles(testItem.obstacleGrid), ShouldEqual, testItem.result)
		})

	}

}
