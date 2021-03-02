package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	metaDebug = 1
	dat := `
	foo = "bar"
`
	err := Parse([]byte(dat))
	assert.NoError(t, err)
}
