package annotation

import (
	"testing"

	"github.com/fader4/okdoc/syntax"
	"github.com/stretchr/testify/assert"
)

func TestPreProcessing(t *testing.T) {
	annotationDebug = 1

	// any symbol
	// @baz
	// @foo(foo = "wqd", z = {"q"}, qwd, qwd )
	// @foobar(
	//     id       = 2868724,
	//     synopsis = "Enable time-travel",
	//     engineer = "Mr. Peabody",
	//     date     = "4/1/3007",
	// 	tags = 	{"a", "b"}
	// )
	dat := `
@CamelCase(bar =
	"b\"a)r",
	baz = "qwdq"
)
qwd

	  	@CamelCase(
		bar = "b\"a)r(",
		(baz) = "qwdq"
	)

qwdqwdqw
@CamelCase(bar = "b\"a)r", baz = "qwdq")
`
	lex, err := newPreprocessing([]byte(dat))
	assert.NoError(t, err)
	lexForParser := &lexerTokenIter{lex: lex, onlyTokens: []string{
		"annotation",
		"bracket",
		"at",
		"ident",
	}}
	var annotations = []syntax.Token{}
	var nextAnnotation syntax.Token
	for {
		container := &annotationSymType{}
		symbol := lexForParser.Lex(container)
		if symbol == 0 {
			break
		}

		switch symbol {
		case beginAnnotation:
			nextAnnotation = syntax.Token{
				Symbol: annotation,
				Start:  container.token.Start,
				Pos:    container.token.Pos,
			}
		case endAnnotation:
			nextAnnotation.End = container.token.End
			nextAnnotation.Pos[1] = nextAnnotation.End - nextAnnotation.Start
			annotations = append(annotations, nextAnnotation)
		}
	}

	t.Log("Num annotations:", len(annotations))

	for _, annot := range annotations {
		lexForParser.PrintToken(&annot)
	}
}
