package annotation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreProcessing(t *testing.T) {
	annotationDebug = 1
	dat := `
any symbol
@baz
@foo(foo = "wqd", z = {"q"}, qwd, qwd )
@foobar(
    id       = 2868724,
    synopsis = "Enable time-travel",
    engineer = "Mr. Peabody",
    date     = "4/1/3007",
	tags = 	{"a", "b"}
)
`
	lex, err := newPreprocessing([]byte(dat))
	assert.NoError(t, err)
	lexForParser := &lexerTokenIter{lex: lex, onlyTokens: []string{
		"annotation",
		"bracket",
		"at",
		"ident",
	}}
	for {
		symbol := lexForParser.Lex(&annotationSymType{})
		if symbol == 0 {
			break
		}
		fmt.Println("TokenSymbol", symbol)
	}
}
