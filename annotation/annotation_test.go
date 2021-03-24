package annotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	annotationDebug = 1
	dat := `
@foo
@foobar()
@foobar(a,b)
@foobar(a.b,b)
@foobar(a='text
\'
text',b)
@foobar(a=1,b)
@foobar(a=1.2,b)
@foobar(a=1.2,b=c.d)
@foobar(a=True,b)
@foobar(
	a=True,
	b)
@foobar(True)
@foobar("True")
@foobar(1)
@foobar(1.2)
@foobar(
    id       = 2868724,
    synopsis = "Enable time-travel",
    engineer = "Mr. Peabody",
    date     = "4/1/3007",
	tags = 	{"a", "b"}
)

`
	err := Parse([]byte(dat))
	assert.NoError(t, err)
}
