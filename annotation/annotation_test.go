package annotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	annotationDebug = 1

	dat := `
any symbol
@Baz
	@Foo(foo = "wqd", z = ["q"], qwd, qwd )
	@Bar
@Foo(foo = "wqd", z = ["q"], qwd, qwd )
	@Foobar(
	id       = 2868724,
	synopsis = "Enable time-travel",
	engineer = "Mr. Peabody",
	date     = "4/1/3007",
	tags = 	["a", "b"]
)

@CamelCase(bar =
	"b\"a)r",
	baz = "qwdq"
)
qwd

	  	@CamelCase(
		bar = "b\"a)r(",
		baz = "qwdq"
	)

qwdqwdqw
@CamelCase(bar = "b\"a)r", baz = "qwdq")
`
	annotations, err := Parse([]byte(dat))
	assert.NoError(t, err)
	t.Log("Num annotations:", len(annotations))

	for _, annot := range annotations {
		annotDat, _ := annot.Token().Bytes([]byte(dat))
		err := parseAnnotation(annot, annotDat)
		assert.NoError(t, err)
		t.Logf(">>%d L%d-L%d(%d): %q\n",
			annot.Start.Pos.Spaces(),
			annot.Start.Pos.Line(),
			annot.End.Pos.Line(),
			annot.Len(), // Size of content of token
			annot.Name,
		)

	}
}

// func TestParse(t *testing.T) {
// 	annotationDebug = 1
// 	dat := `
// @foo
//   @foobar()
// @foobar(a,b)
// @foobar(a.b,b)
// @foobar(a='text
// \'
// text',b)
// @foobar(a=1,b)
// @foobar(a=1.2,b)
// @foobar(a=1.2,b=c.d)
// @foobar(a=True,b)
// @foobar(
// 	a=True,
// 	b)
// @foobar(True)
// @foobar("True")
// @foobar(1)
// @foobar(1.2)
// @foobar(
//     id       = 2868724,
//     synopsis = "Enable time-travel",
//     engineer = "Mr. Peabody",
//     date     = "4/1/3007",
// 	tags = 	["a", "b"],
// 	obj = {
// 		foo = "bar",
// 		arr = ["a", "b"],
// 		baz = {
// 			foo = "bar",
// 			baz = baz,
// 			arr = ["a", "b"],
// 			arr = [{}, {b = "b"}],
// 			arr = [[1,2], [2,3,4]]
// 		}
// 	}
// )

// `
// 	_, err := Parse([]byte(dat))
// 	assert.NoError(t, err)
// }
