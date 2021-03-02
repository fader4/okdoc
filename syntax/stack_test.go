package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelper_IntStack(t *testing.T) {
	s := IntStack{}
	assert.EqualValues(t, -1, s.Top())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.EqualValues(t, 3, s.Top())
	assert.EqualValues(t, 3, s.Pop())

	assert.EqualValues(t, 2, s.Top())
}
