package starlark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerHelper_stack(t *testing.T) {
	s := stackChars{}
	assert.EqualValues(t, -1, s.top())
	s.push(1)
	s.push(2)
	s.push(3)
	assert.EqualValues(t, 3, s.top())
	assert.EqualValues(t, 3, s.pop())

	assert.EqualValues(t, 2, s.top())
}
