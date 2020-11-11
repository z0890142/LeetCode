package EvaluateReversePolishNotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evalRPN(t *testing.T) {
	tests := []struct {
		tokens []string
		result int
	}{
		{
			[]string{"2", "1", "+", "3", "*"},
			9,
		},
		{
			[]string{"4", "13", "5", "/", "+"},
			6,
		},
		{
			[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			22,
		},
	}

	for _, testItem := range tests {
		assert.Equal(t, testItem.result, evalRPN(testItem.tokens))
	}

}
