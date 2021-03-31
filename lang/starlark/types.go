package starlark

import "github.com/fader4/okdoc/syntax"

type Comment struct {
	*syntax.TokenWithData
	Multiline bool
}

type Return struct {
	*syntax.TokenWithData
}

type Def struct {
	*syntax.TokenWithData

	Name   syntax.Value
	Fields []*DefField
}

type DefField struct {
	Key     syntax.Value
	Value   syntax.Value
	Varargs bool
	Kwargs  bool
}

type Load struct {
	*syntax.TokenWithData

	Fields syntax.Array
}

type Module struct {
	*syntax.TokenWithData

	Ident  syntax.Ident_
	Name   syntax.StringLiteral
	Fields syntax.Array
}

type CallFunc struct {
	Name               syntax.Ident_
	Fields             syntax.Array
	ArbitraryArrArgs   syntax.Array
	ArbitraryNamedArgs []*DictField
}

func (f CallFunc) FuncName() string {
	// без @
	return f.Name.String()[1:]
}

type DictField struct {
	Key         syntax.Value
	Value       syntax.Value
	ValueDict   syntax.Value
	ValueCallFn *CallFunc
}
