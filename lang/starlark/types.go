package starlark

import (
	"bytes"
	"fmt"

	"github.com/fader4/okdoc/syntax"
)

type Comment struct {
	*syntax.TokenWithData
	Multiline bool
}

type Return struct {
	*syntax.TokenWithData
}

type Def struct {
	*syntax.TokenWithData

	Name   syntax.Ident_
	Fields []*DefField
}

type DefField struct {
	Key     *Operand
	Value   *Operand
	Varargs bool
	Kwargs  bool
}

type Load struct {
	*syntax.TokenWithData

	Fields syntax.Array
}

type Module struct {
	*syntax.TokenWithData

	Export syntax.Ident_
	Name   syntax.StringLiteral
	Fields syntax.Array
}

type CallFunc struct {
	Name syntax.Ident_
	// free fields => ((key1, val1), (key2, null), ...)
	// *vararrgs => (val1, val2, ...)
	// *kwargs => ((key1, val1), (key2, val2), ...)
	Fields syntax.Array

	Varrarrgs bool
	Kwarrgs   bool
}

func (f CallFunc) FuncName() string {
	// без @
	return f.Name.String()[1:]
}

type Operand struct {
	Parent   *Operand
	Value    syntax.Value
	CallFunc *CallFunc
}

func (o *Operand) String() string {
	var arr = []syntax.Value{}

	buf := &bytes.Buffer{}
	walk := func(buf *bytes.Buffer, o *Operand) (next *Operand) {
		if o.CallFunc != nil {
			buf.WriteString(".Call(" + o.CallFunc.FuncName() + ")")
		}
		if o.Value != nil {
			fmt.Fprintf(buf, ".%v", o.Value)
		}
		return o.Parent
	}
	next := o
	for {
		arr = append(arr, next)
		next = walk(buf, next)
		if next == nil {
			break
		}
	}
	return buf.String()[1:]
}

type Assign struct {
	Left  *Operand
	Op    int
	Right *Operand
}
