package SimplifyPath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_simplifyPath(t *testing.T) {
	tests := []struct {
		path   string
		result string
	}{
		{
			"/home/",
			"/home",
		},
		{
			"/../",
			"/",
		},
		{

			"/a//b////c/d//././/..",
			"/a/b/c",
		},
		{
			"/a/./b/../../c/",
			"/c",
		}, {
			"/a/../../b/../c//.//",
			"/c",
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, simplifyPath(testItem.path))

	}
}
