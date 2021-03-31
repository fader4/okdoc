package starlark

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	v1 := `
load(
	"foo",
	bar = "baz"
)

# inline comment

foo = module(
	"bar"
)

"""
Free comment
"""

def main(
	ctx=(1,2,3)
	):
	"""
	Inline comment into method
	"""

	if False:
		# inline comment
		return False

	return True
`
	res, err := Parse([]byte(v1))
	assert.NoError(t, err)
	fmt.Println("---- list tokens ----")
	for _, token := range res {
		str, err := token.HumanString()
		assert.NoError(t, err)
		fmt.Println(str)
	}
}
