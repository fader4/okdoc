package starlark

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	v1 := `
load("a", b = "c", "qw")

# inline comment

foo = module(
	"bar"
)

foo = module(
	"bar",
	a = "b",
	arr = 1,
	ar = f(**dict(c=7, obj=dict(c=7, a=2, b=3), a=2, b=3, fn=f(**dict(c=7, a=2, b=3)))),
	dic = dict(c=7, a=2, b=3)
)

"""
Free comment
"""

def main(ctx=(1,2,3) , **kwargs):
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
		err = token.Parse()
		assert.NoError(t, err)
	}
}
