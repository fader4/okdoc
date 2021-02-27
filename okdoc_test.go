package okdoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ = ""
var case1 = `
return ("bar")
	"""
	foo bar
	foo # not comment

	"""

def main (foo, bar = "bar", **kwars) :
	"foo bar"

	def main(foo, bar = "bar", **kwars) :
		foo

    def main(foo, bar = "bar", **kwars) :
		"""
		some comments
		"""
def main(foo, bar = "bar", **kwars) :
	foo
`

func TestParse(t *testing.T) {
	yyDebug = 1
	yyErrorVerbose = true
	res, err := Parse([]byte(case1))
	assert.NoError(t, err)
	for _, n := range res.Nodes {
		t.Log("\t", n)
	}
}
