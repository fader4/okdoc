package starlark

type ArgumentNode []ArgumentNode_Value

type ArgumentNode_Value struct {
	Ident         string
	StringLiteral string
}
