package PopulatingNextRightPointersinEachNode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connect(t *testing.T) {
	tests := []struct {
		root   *Node
		result *Node
	}{
		{
			&Node{
				1,
				&Node{
					2, nil, nil, nil,
				},
				&Node{
					3, nil, nil, nil,
				},
				nil,
			},
			&Node{
				1,
				&Node{
					2, nil, nil, nil,
				},
				&Node{
					3, nil, nil, nil,
				},
				nil,
			},
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, connect(testItem.root))

	}
}
