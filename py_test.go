package okdoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var case1 = `

load ("foo", baz = "baz")

def main(
	ctx = "bar",
	baz,
	**kwars):
	"comment for main function"
	return "ok"

# 1 inline comments
	# 2 inline comments with tabs

foo = "bar" # 3 inline comments
foo, bar = 1, 2

"""
qwd
"""
`

func TestPyParse(t *testing.T) {
	pyDebug = 1
	pyErrorVerbose = true

	err := PyParse([]byte(case1))
	assert.NoError(t, err)
}
