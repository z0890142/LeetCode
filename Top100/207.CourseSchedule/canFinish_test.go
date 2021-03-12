package CourseSchedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canFinish(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		result        bool
	}{
		{
			2,
			[][]int{[]int{1, 0}},
			true,
		},
	}
	for _, testItem := range tests {
		assert.Equal(t, testItem.result, canFinish(testItem.numCourses, testItem.prerequisites))
	}

}
